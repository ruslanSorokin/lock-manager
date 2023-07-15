package service

import (
	"context"
	"errors"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/provider"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

func (s LockService) Unlock(
	ctx context.Context,
	rID, tkn string,
) error {
	switch {
	case !isValidResourceID(rID):
		return ErrInvalidResourceID

	case !isValidToken(tkn):
		return ErrInvalidToken

	default:
	}

	err := s.lockProvider.DeleteIfTokenMatches(
		ctx,
		model.NewLock(
			rID,
			tkn,
		),
	)

	switch {
	case errors.Is(err, provider.ErrLockNotFound) || errors.Is(err, provider.ErrWrongToken):
		s.log.Info(
			err.Error(),
			"resourceID", rID,
			"token", tkn,
		)
		return Errf(err)

	case err != nil:
		s.log.Error(
			err,
			"resourceID", rID,
			"token", tkn,
		)
		return Errf(err)

	default:
		return nil
	}
}
