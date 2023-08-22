package service

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

func (s LockService) Lock(
	ctx context.Context,
	rID string,
) (string, error) {
	if err := s.validateResourceID(rID); err != nil {
		return "", err
	}

	l := model.NewLockWithToken(rID)

	err := s.lockProvider.Create(ctx, l)
	if err != nil {
		return "", Errf(err)
	}

	s.mtr.IncLockedTotal()
	return l.Token, nil
}
