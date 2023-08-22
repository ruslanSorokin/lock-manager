package provider

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

//go:generate go run github.com/vektra/mockery/v2@latest --name=LockProviderI --structname LockProvider --output=mock --case=underscore --disable-version-string --outpkg=mock

//go:generate go run mvdan.cc/gofumpt@latest -l -w mock/lock_provider_i.go

// LockProviderI provides CRUD+Custom operations on the `Lock` model.
//
// Uniqueness is guaranteed by `Lock.ResourceID` field.
type LockProviderI interface {
	// Create creates a new `lock`.
	//
	// If `lock` already exists(by `lock`.ResourceID), then `ErrLockAlreadyExists`
	// is returned.
	Create(ctx context.Context, lock model.Lock) error

	// Delete deletes lock with given `resourceID`.
	//
	// If there is no such lock(by `lock`.ResourceID), then `ErrLockNotFound` is
	// returned.
	Delete(ctx context.Context, resourceID string) error

	// Get returns lock with given `resourceID`.
	//
	// If there is no such lock(by `lock`.ResourceID), then `ErrLockNotFound` is
	// returned.
	Get(ctx context.Context, resourceID string) (model.Lock, error)

	// DeleteIfTokenMatches deletes lock with given `lock`.ResourceID only if
	// `lock`.Token fits.
	//
	// If there is no such `lock`(by `lock`.ResourceID), then
	// `ErrLockNotFound` is returned.
	//
	// If `lock`.Token doesn't fit, then `ErrWrongToken` is returned.
	DeleteIfTokenMatches(ctx context.Context, lock model.Lock) error
}
