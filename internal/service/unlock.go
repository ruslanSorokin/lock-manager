package service

import (
	"context"
	"errors"

	"github.com/ruslanSorokin/lock-manager/internal/model"
)

func (s LockService) Unlock(
	ctx context.Context,
	rID, tkn string,
) error {
	var errs []error
	if err := s.validateResourceID(rID); err != nil {
		errs = append(errs, Errf(err))
	}
	if err := s.validateToken(tkn); err != nil {
		errs = append(errs, Errf(err))
	}
	if len(errs) != 0 {
		return errors.Join(errs...)
	}

	err := s.lockProvider.DeleteIfTokenMatches(
		ctx,
		model.NewLock(
			rID, tkn,
		),
	)
	if err != nil {
		return Errf(err)
	}

	return nil
}
