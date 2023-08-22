package provider

import (
	"context"
	"testing"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testDeleteIfTokenMatches(t *testing.T, s LockProviderI) {
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

		err = s.DeleteIfTokenMatches(
			ctx,
			tc.l,
		)
		require.NoError(err,
			"must delete the lock without any error",
		)

		_, err = s.Get(ctx, tc.l.ResourceID)
		assert.ErrorIsf(err, ErrLockNotFound,
			"must return %w, as we've just deleted the lock",
			ErrLockNotFound,
		)
	}
}

func testDeleteIfTokenMatchesErrInvalidToken(t *testing.T, s LockProviderI) {
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

		wrongToken := "wrong token here"
		require.NotEqual(tc.l.Token, wrongToken,
			"must be not the same as the right one",
		)

		err = s.DeleteIfTokenMatches(
			ctx,
			model.NewLock(
				tc.l.ResourceID,
				wrongToken,
			),
		)
		require.ErrorIsf(err, ErrWrongToken,
			"must return %w, as we use wrong token for deletion",
			ErrWrongToken,
		)
	}
}

func testDeleteIfTokenMatchesErrLockNotFound(t *testing.T, s LockProviderI) {
	require := require.New(t)
	assert := assert.New(t)

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

		err := s.DeleteIfTokenMatches(
			ctx,
			tc.l,
		)
		require.ErrorIsf(err, ErrLockNotFound,
			"must return %w, as there is no such lock in the storage",
			ErrLockNotFound,
		)

		err = s.Create(ctx, tc.l)
		assert.NoError(err,
			"should create lock without any error",
		)

		err = s.DeleteIfTokenMatches(
			ctx,
			tc.l,
		)
		require.NoError(err, ErrLockNotFound,
			"must delete lock without any error",
			ErrLockNotFound,
		)

		err = s.DeleteIfTokenMatches(
			ctx,
			tc.l,
		)
		require.ErrorIsf(err, ErrLockNotFound,
			"must return %w, as there is no such lock in the storage",
			ErrLockNotFound,
		)
	}
}
