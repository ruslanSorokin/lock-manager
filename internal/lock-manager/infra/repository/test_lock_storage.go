package repository

import (
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func getTestName[T any](f T) string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	_, a, ok := strings.Cut(name, ".test")
	if !ok {
		panic("wrong package name")
	}
	return a
}

// RunLockStorageTests runs each test on the LockStorage, received from ctor.
// Before each test and in the end, storage gets flushed with the flusher.
func RunLockStorageTests(t *testing.T,
	ctor func() LockStorageI, flusher func() error) {
	require := require.New(t)

	// TODO: use reflect for this purpose
	// see: https://github.com/ungerik/pkgreflect
	tests := []func(*testing.T, LockStorageI){
		testCreate,
		testCreateErrLockAlreadyExists,

		testDelete,
		testDeleteErrLockNotFound,

		testDeleteIfTokenMatches,
		testDeleteIfTokenMatchesErrInvalidToken,
		testDeleteIfTokenMatchesErrLockNotFound,

		testGet,
		testGetErrLockNotFound,
	}

	for _, test := range tests {
		storage := ctor()
		err := flusher()
		require.NoError(err,
			"unable to flush storage",
		)

		t.Run(getTestName(test), func(t *testing.T) {
			test(t, storage)
		})
	}

	err := flusher()
	require.NoError(err,
		"unable to flush storage",
	)
}
