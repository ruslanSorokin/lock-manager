package iredis

import (
	"context"
	"time"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

func (s LockStorage) Create(ctx context.Context, l *model.Lock) error {
	didSet, err := s.conn.DB.SetNX(ctx, l.ResourceID(), l.Token(), 0*time.Second).
		Result()
	if err != nil {
		s.l.Error(err,
			"lock", l)
		return Errf(err)
	}
	if !didSet {
		err = ErrLockAlreadyExists
		return Errf(err)
	}

	return nil
}
