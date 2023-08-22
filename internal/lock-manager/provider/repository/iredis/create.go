package iredis

import (
	"context"
	"time"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
)

func (s LockStorage) Create(ctx context.Context, l model.Lock) error {
	didSet, err := s.conn.DB.SetNX(ctx, l.ResourceID, l.Token, 0*time.Second).Result()
	if err != nil {
		s.l.Error(
			err,
			"lock", l,
		)
		return provider.Errf(err)
	}
	if !didSet {
		err = provider.ErrLockAlreadyExists
		return provider.Errf(err)
	}

	return nil
}