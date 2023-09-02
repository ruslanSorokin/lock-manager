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
		return "", Errorf(err)
	}

	l, err := model.NewLockWithToken(rID)
	if err != nil {
		return "", err
	}

	if err = s.lockProvider.Create(ctx, l); err != nil {
		return "", Errorf(err)
	}

	s.mtr.IncLockedTotal()
	return l.Token(), nil
}
