package provider

import (
	"context"
	"testing"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testDelete(t *testing.T, s LockProviderI) {
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
			"should create the lock without any error",
		)

		err = s.Delete(ctx, tc.l.ResourceID)
		require.NoError(err,
			"must return the lock without any error",
		)

		_, err = s.Get(ctx, tc.l.ResourceID)
		assert.ErrorIsf(err, ErrLockNotFound,
			"must return %w, as we've just deleted the lock",
			ErrLockNotFound,
		)
	}
}

func testDeleteErrLockNotFound(t *testing.T, s LockProviderI) {
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

		err := s.Delete(ctx, tc.l.ResourceID)
		require.ErrorIsf(err, ErrLockNotFound,
			"must return %w, as there is no such lock in the storage",
			ErrLockNotFound,
		)
	}
}
