package iredis

import (
	"context"
	"errors"

	redis "github.com/redis/go-redis/v9"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

func (s LockStorage) Get(
	ctx context.Context,
	resourceID string,
) (*model.Lock, error) {
	res, err := s.conn.DB.Get(ctx, resourceID).Result()
	if errors.Is(err, redis.Nil) {
		err = ErrLockNotFound
		return nil, Errf(err)
	}
	if err != nil {
		s.l.Error(err,
			"resourceID", resourceID,
		)
		return nil, Errf(err)
	}

	l, err := model.NewLock(resourceID, res)
	if err != nil {
		return nil, err
	}
	return l, nil
}
