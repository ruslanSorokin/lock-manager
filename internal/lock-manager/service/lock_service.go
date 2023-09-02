package service

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/go-playground/validator/v10"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/imetric"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
)

type LockServiceI interface {
	// Lock locks given `resourceID` and returns a token.
	//
	// If given `resourceID` is not valid, then `service.ErrInvalidResourceID` is
	// returned.
	//
	// If given `resourceID` is already locked, then
	// `provider.ErrLockAlreadyExists` is returned.
	Lock(ctx context.Context, resourceID string) (string, error)

	// Unlock unlocks given `resourceID` only if `token` fits.
	//
	// If given `resourceID` is not valid, then `service.ErrInvalidResourceID` is
	// returned.
	//
	// If given `token` is not valid, then `service.ErrInvalidToken` is returned.
	//
	// If there is no such lock, then `provider.ErrLockNotFound` is returned.
	//
	// If given `token` doesn't fit, then `provider.ErrWrongToken` is returned.
	Unlock(ctx context.Context, resourceID, token string) error
}

type LockService struct {
	log       logr.Logger
	validator *validator.Validate

	mtr          imetric.ServiceMetricI
	lockProvider provider.LockProviderI

	resourceIDValidator resourceIDValidator
	tokenValidator      tokenValidator
}

var _ LockServiceI = (*LockService)(nil)

func New(
	l logr.Logger,
	v *validator.Validate,
	lp provider.LockProviderI,
	m imetric.ServiceMetricI,
) *LockService {
	return &LockService{
		log:                 l,
		validator:           v,
		mtr:                 m,
		lockProvider:        lp,
		resourceIDValidator: newResourceIDValidator(v),
		tokenValidator:      newTokenValidator(v),
	}
}
