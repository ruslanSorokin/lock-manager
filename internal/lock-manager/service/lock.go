package service

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

func (s LockService) Lock(
	ctx context.Context,
	rID string,
) (string, error) {
	if err := s.resourceIDValidator(rID); err != nil {
		return "", Errf(err)
	}

	l, err := model.NewLock(rID)
	if err != nil {
		return "", err
	}

	if err = s.lockProvider.Create(ctx, l); err != nil {
		return "", Errf(err)
	}

	s.mtr.IncLockedTotal()
	return l.Token(), nil
}
