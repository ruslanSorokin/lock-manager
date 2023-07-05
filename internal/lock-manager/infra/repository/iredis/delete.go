package iredis

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/repository"
)

func (s LockStorage) Delete(ctx context.Context, resourceID string) error {
	delCount, err := s.db.Del(ctx, resourceID).Result()
	if err != nil {
		s.l.Error(
			err,
			"resourceID", resourceID,
		)

		return repository.Errf(err)
	}
	if delCount != 1 {
		err = repository.ErrLockNotFound
		s.l.Info(
			err.Error(),
			"resourceID", resourceID,
		)

		return repository.Errf(err)
	}

	return nil
}
