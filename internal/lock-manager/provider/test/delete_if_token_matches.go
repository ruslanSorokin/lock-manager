package providertest

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
)

func (s PSuite) TestDeleteIfTokenMatches() {
	t := s.T()
	p := s.Provider
	assert := assert.New(t)
	require := require.New(t)

	tcs := []struct {
		l *model.Lock
	}{
		{
			l: Must(model.NewLock(
				"path/to/resource",
				uuid.Must(uuid.NewV4()).String(),
			)),
		},
		{
			l: Must(model.NewLock(
				"another/path/to/resource",
				uuid.Must(uuid.NewV4()).String(),
			)),
		},
	}
	for _, tc := range tcs {
		ctx := context.Background()

		err := p.Create(ctx, tc.l)
		assert.NoError(err,
			"should create the lock without any error",
		)

		err = p.DeleteIfTokenMatches(
			ctx,
			tc.l,
		)
		require.NoError(err,
			"must delete the lock without any error",
		)

		_, err = p.Get(ctx, tc.l.ResourceID())
		assert.ErrorIsf(err, provider.ErrLockNotFound,
			"must return %w, as we've just deleted the lock",
			provider.ErrLockNotFound,
		)
	}
}

func (s PSuite) TestDeleteIfTokenMatchesErrInvalidToken() {
	t := s.T()
	p := s.Provider
	assert := assert.New(t)
	require := require.New(t)

	tcs := []struct {
		l *model.Lock
	}{
		{
			l: Must(model.NewLock(
				"path/to/resource",
				uuid.Must(uuid.NewV4()).String(),
			)),
		},
		{
			l: Must(model.NewLock(
				"another/path/to/resource",
				uuid.Must(uuid.NewV4()).String(),
			)),
		},
	}
	for _, tc := range tcs {
		ctx := context.Background()

		err := p.Create(ctx, tc.l)
		assert.NoError(err,
			"should create lock without any error",
		)

		wrongToken := uuid.Must(uuid.NewV4()).String()
		require.NotEqual(tc.l.Token(), wrongToken,
			"must be not the same as the right one",
		)

		err = p.DeleteIfTokenMatches(
			ctx,
			Must(model.NewLock(
				tc.l.ResourceID(),
				wrongToken,
			)),
		)
		require.ErrorIsf(err, provider.ErrWrongToken,
			"must return %w, as we use wrong token for deletion",
			provider.ErrWrongToken,
		)
	}
}

func (s PSuite) TestDeleteIfTokenMatchesErrLockNotFound() {
	t := s.T()
	p := s.Provider
	require := require.New(t)
	assert := assert.New(t)

	tcs := []struct {
		l *model.Lock
	}{
		{
			l: Must(model.NewLock(
				"path/to/resource",
				uuid.Must(uuid.NewV4()).String(),
			)),
		},
		{
			l: Must(model.NewLock(
				"another/path/to/resource",
				uuid.Must(uuid.NewV4()).String(),
			)),
		},
	}
	for _, tc := range tcs {
		ctx := context.Background()

		err := p.DeleteIfTokenMatches(
			ctx,
			tc.l,
		)
		require.ErrorIsf(err, provider.ErrLockNotFound,
			"must return %w, as there is no such lock in the storage",
			provider.ErrLockNotFound,
		)

		err = p.Create(ctx, tc.l)
		assert.NoError(err,
			"should create lock without any error",
		)

		err = p.DeleteIfTokenMatches(
			ctx,
			tc.l,
		)
		require.NoError(err, provider.ErrLockNotFound,
			"must delete lock without any error",
			provider.ErrLockNotFound,
		)

		err = p.DeleteIfTokenMatches(
			ctx,
			tc.l,
		)
		require.ErrorIsf(err, provider.ErrLockNotFound,
			"must return %w, as there is no such lock in the storage",
			provider.ErrLockNotFound,
		)
	}
}
