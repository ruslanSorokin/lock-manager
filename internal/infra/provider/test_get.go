package provider

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/lock-manager/internal/model"
)

func testGet(t *testing.T, s LockProviderI) {
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
		assert.NoError(err,
			"should create lock without any error",
		)

		l, err := s.Get(ctx, tc.l.ResourceID)
		require.NoError(err,
			"must return lock without any error",
		)
		require.Equal(l, tc.l,
			"must be the same, as it was before inserting into the storage",
		)
	}
}

func testGetErrLockNotFound(t *testing.T, s LockProviderI) {
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

		_, err := s.Get(ctx, tc.l.ResourceID)
		require.ErrorIsf(err, ErrLockNotFound,
			"must return %w, as there is no such lock in the storage",
			ErrLockNotFound,
		)
	}
}
