package iredis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/repository"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

func (s LockStorage) Get(ctx context.Context, resourceID string) (model.Lock, error) {
	res, err := s.db.Get(ctx, resourceID).Result()
	if errors.Is(err, redis.Nil) {
		err = repository.ErrLockNotFound
		s.l.Info(
			err.Error(),
			"lock", resourceID,
		)

		return model.Lock{}, repository.Errf(err)
	}
	if err != nil {
		s.l.Error(
			err,
			"resourceID", resourceID,
		)

		return model.Lock{}, repository.Errf(err)
	}

	return model.Lock{
		ResourceID: resourceID,
		Token:      res,
	}, nil
}
