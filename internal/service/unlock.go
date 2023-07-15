package service

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/model"
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
	if err != nil {
		return Errf(err)
	}

	return nil
}
