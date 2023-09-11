package iredis

import (
	"context"
)

func (s LockStorage) Delete(ctx context.Context, resourceID string) error {
	delCount, err := s.conn.DB.Del(ctx, resourceID).Result()
	if err != nil {
		s.l.Error(
			err,
			"resourceID", resourceID,
		)
		return Errf(err)
	}
	if delCount != 1 {
		err = ErrLockNotFound
		return Errf(err)
	}

	return nil
}
