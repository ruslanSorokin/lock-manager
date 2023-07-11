package provider

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=LockProviderI --structname LockProvider --output=mock --case=underscore --disable-version-string --outpkg=mock

//go:generate go run mvdan.cc/gofumpt@latest -l -w mock/lock_provider_i.go

// LockProviderI provides CRUD+Custom operations on the lock model.
type LockProviderI interface {
	// Create creates a new lock.
	//
	// If lock with such resourceID already exists, then ErrLockAlreadyExists is
	// returned.
	Create(ctx context.Context, l model.Lock) error

	// Delete deletes lock with given resourceID.
	//
	// If there is no such lock, then ErrLockNotFound is returned.
	Delete(ctx context.Context, resourceID string) error

	// Get returns lock with given resourceID.
	//
	// If there is no such lock, then ErrLockNotFound is returned.
	Get(ctx context.Context, resourceID string) (model.Lock, error)

	// DeleteIfTokenMatches deletes lock with given resourceID only if the token
	// is the same as the one that already in the storage.
	//
	// If there is no such lock, then ErrLockNotFound is returned.
	//
	// If token is not the same as the one that is already in the storage, then
	// ErrInvalidToken is returned.
	DeleteIfTokenMatches(ctx context.Context, lock model.Lock) error
}
