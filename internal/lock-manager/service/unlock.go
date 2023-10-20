package service

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

func (s LockService) Unlock(
	ctx context.Context,
	rID string,
	tkn string,
) error {
	if err := s.resourceIDValidator(rID); err != nil {
		return Errf(err)
	}
	if err := s.tokenValidator(tkn); err != nil {
		return Errf(err)
	}

	l, err := model.ReinstateLock(rID, tkn)
	if err != nil {
		return err
	}

	if err := s.lockProvider.DeleteIfTokenMatches(ctx, l); err != nil {
		return Errf(err)
	}

	s.mtr.IncUnlockedTotal()
	return nil
}
