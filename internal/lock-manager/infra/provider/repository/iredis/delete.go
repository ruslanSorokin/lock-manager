package iredis

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/provider"
)

func (s LockStorage) Delete(ctx context.Context, resourceID string) error {
	delCount, err := s.db.Del(ctx, resourceID).Result()
	if err != nil {
		s.l.Error(
			err,
			"resourceID", resourceID,
		)
		return provider.Errf(err)
	}
	if delCount != 1 {
		err = provider.ErrLockNotFound
		s.l.Info(
			err.Error(),
			"resourceID", resourceID,
		)
		return provider.Errf(err)
	}

	return nil
}
