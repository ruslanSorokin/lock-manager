package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

func testCreate(t *testing.T, s LockStorageI) {
	assert := assert.New(t)
	require := require.New(t)

	tcs := []struct {
		l model.Lock
	}{
		{
			l: model.NewLock(
				"path/to/resource",
				"token12345",
			),
		},
		{
			l: model.NewLock(
				"another/path/to/resource",
				"token1234567890",
			),
		},
	}
	for _, tc := range tcs {
		ctx := context.Background()

		err := s.Create(ctx, tc.l)
		require.NoError(err,
			"should create lock without any error",
		)

		l, err := s.Get(ctx, tc.l.ResourceID)
		assert.NoError(err,
			"must return lock without any error",
		)
		require.Equal(l, tc.l,
			"must be the same, as it was before inserting into the storage",
		)
	}
}

func testCreateErrLockAlreadyExists(t *testing.T, s LockStorageI) {
	require := require.New(t)

	tcs := []struct {
		l model.Lock
	}{
		{
			l: model.NewLock(
				"path/to/resource",
				"token12345",
			),
		},
		{
			l: model.NewLock(
				"another/path/to/resource",
				"token1234567890",
			),
		},
	}
	for _, tc := range tcs {
		ctx := context.Background()

		err := s.Create(ctx, tc.l)
		require.NoError(err,
			"must insert without any error, as there is no such lock in the storage",
		)

		err = s.Create(ctx, tc.l)
		require.ErrorIsf(err, ErrLockAlreadyExists,
			"must return %w, as there is already such lock in the storage",
			ErrLockAlreadyExists,
		)
	}
}
