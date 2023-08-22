package iredis

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
)

func (s LockStorage) Delete(ctx context.Context, resourceID string) error {
	delCount, err := s.conn.DB.Del(ctx, resourceID).Result()
	if err != nil {
		s.l.Error(
			err,
			"resourceID", resourceID,
		)
		return provider.Errf(err)
	}
	if delCount != 1 {
		err = provider.ErrLockNotFound
		return provider.Errf(err)
	}

	return nil
}
