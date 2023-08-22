package service

import (
	"context"

	"github.com/go-logr/logr"
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
	mtr          imetric.ServiceMetricI
	log          logr.Logger
	lockProvider provider.LockProviderI

	validateResourceID resourceIDValidator
	validateToken      tokenValidator
}

var _ LockServiceI = (*LockService)(nil)

func New(
	l logr.Logger,
	lp provider.LockProviderI,
	m imetric.ServiceMetricI,
	rIDdMinLen, rIDMaxLen int,
) *LockService {
	return &LockService{
		mtr:                m,
		log:                l,
		lockProvider:       lp,
		validateResourceID: newResourceIDValidator(rIDdMinLen, rIDMaxLen),
		validateToken:      newTokenValidator(),
	}
}

func NewFromConfig(
	l logr.Logger,
	lp provider.LockProviderI,
	m imetric.ServiceMetricI,
	cfg *Config,
) *LockService {
	return New(
		l, lp, m,
		cfg.ResourceID.MinLen, cfg.ResourceID.MaxLen)
}

func Default(
	l logr.Logger,
	lp provider.LockProviderI,
	m imetric.ServiceMetricI,
) *LockService {
	return New(
		l, lp, m,
		defaultResourceIDMinLen, defaultResourceIDMaxLen,
	)
}
