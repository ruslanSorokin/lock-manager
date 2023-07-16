package service

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/model"
)

func (s LockService) Lock(
	ctx context.Context,
	rID string,
) (string, error) {
	if !s.isValidResourceID(rID) {
		return "", ErrInvalidResourceID
	}

	l := model.NewLockWithToken(rID)

	err := s.lockProvider.Create(ctx, l)
	if err != nil {
		return "", Errf(err)
	}

	return l.Token, nil
}
