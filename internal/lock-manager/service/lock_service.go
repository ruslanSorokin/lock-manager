package service

import (
	"context"

	"github.com/go-logr/logr"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/provider"
)

type LockServiceI interface {
	Lock(ctx context.Context, resourceID string) (string, error)
	Unlock(ctx context.Context, resourceID, token string) error
}

type LockService struct {
	log          logr.Logger
	lockProvider provider.LockProviderI
}

var _ LockServiceI = (*LockService)(nil)

func NewLockService(l logr.Logger, lp provider.LockProviderI) LockService {
	return LockService{
		log:          l,
		lockProvider: lp,
	}
}
