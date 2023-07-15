package service

import (
	"context"
	"errors"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/provider"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

func (s LockService) Lock(
	ctx context.Context,
	rID string,
) (string, error) {
	if !isValidResourceID(rID) {
		return "", ErrInvalidResourceID
	}

	l := model.NewLockWithToken(rID)

	err := s.lockProvider.Create(ctx, l)
	switch {
	case errors.Is(err, provider.ErrLockAlreadyExists):
		s.log.Info(
			err.Error(),
			"resourceID", rID,
		)
		return "", Errf(err)

	case err != nil:
		s.log.Error(
			err,
			"resourceID", rID,
		)
		return "", Errf(err)

	default:
		return l.Token, nil
	}
}
