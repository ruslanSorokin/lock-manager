package iredis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
)

func (s LockStorage) Get(ctx context.Context, resourceID string) (model.Lock, error) {
	res, err := s.conn.DB.Get(ctx, resourceID).Result()
	if errors.Is(err, redis.Nil) {
		err = provider.ErrLockNotFound
		return model.Lock{}, provider.Errf(err)
	}
	if err != nil {
		s.l.Error(err,
			"resourceID", resourceID,
		)
		return model.Lock{}, provider.Errf(err)
	}

	return model.Lock{
		ResourceID: resourceID,
		Token:      res,
	}, nil
}
