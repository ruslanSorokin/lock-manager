package service

import (
	"context"

	"github.com/go-logr/logr"

	"github.com/ruslanSorokin/lock-manager/internal/infra/provider"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=LockServiceI --structname LockService --output=mock --case=underscore --disable-version-string --outpkg=mock

//go:generate go run mvdan.cc/gofumpt@latest -l -w mock/lock_service_i.go

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
	log          logr.Logger
	lockProvider provider.LockProviderI

	validateResourceID resourceIDValidator
	validateToken      tokenValidator
}

var _ LockServiceI = (*LockService)(nil)

func NewLockService(
	l logr.Logger,
	lp provider.LockProviderI,
	rIDdMinLen, rIDMaxLen int,
	tknMinLen, tknMaxLen int,
) LockService {
	return LockService{
		log:                l,
		lockProvider:       lp,
		validateResourceID: newResourceIDValidator(rIDdMinLen, rIDMaxLen),
		validateToken:      newTokenValidator(tknMinLen, tknMaxLen),
	}
}

func NewLockServiceFromConfig(
	l logr.Logger,
	lp provider.LockProviderI,
	cfg Config,
) LockService {
	return NewLockService(
		l, lp,
		cfg.ResourceID.MinLen, cfg.ResourceID.MaxLen,
		cfg.Token.MinLen, cfg.Token.MaxLen,
	)
}

func NewLockServiceWithDefaults(
	l logr.Logger,
	lp provider.LockProviderI,
) LockService {
	return NewLockService(
		l, lp,
		defaultResourceIDMinLen, defaultResourceIDMaxLen,
		defaultTokenMinLen, defaultTokenMaxLen,
	)
}