package providertest

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
)

func (s PSuite) TestDelete() {
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

		err = p.Delete(ctx, tc.l.ResourceID())
		require.NoError(err,
			"must return the lock without any error",
		)

		_, err = p.Get(ctx, tc.l.ResourceID())
		assert.ErrorIsf(err, provider.ErrLockNotFound,
			"must return %w, as we've just deleted the lock",
			provider.ErrLockNotFound,
		)
	}
}

func (s PSuite) TestDeleteErrLockNotFound() {
	t := s.T()
	p := s.Provider
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

		err := p.Delete(ctx, tc.l.ResourceID())
		require.ErrorIsf(err, provider.ErrLockNotFound,
			"must return %w, as there is no such lock in the storage",
			provider.ErrLockNotFound,
		)
	}
}
