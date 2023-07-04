package iredis

import (
	"context"
	"time"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/repository"
)

func (s LockStorage) Create(ctx context.Context, l model.Lock) error {
	didSet, err := s.db.SetNX(ctx, l.ResourceID, l.Token, 0*time.Second).Result()
	if err != nil {
		s.l.Error(
			err,
			"lock", l,
		)

		return err
	}
	if !didSet {
		err = repository.ErrLockAlreadyExists
		s.l.Info(
			err.Error(),
			"lock", l,
		)

		return err
	}

	return nil
}
