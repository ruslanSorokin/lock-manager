package provider

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

// LockProviderI provides CRUD+Custom operations on the `Lock` model.
//
// Uniqueness is guaranteed by `Lock.ResourceID` field.
type LockProviderI interface {
	// Create creates a new `lock`.
	//
	// If `lock` already exists(by `lock`.ResourceID), then `ErrLockAlreadyExists`
	// is returned.
	Create(ctx context.Context, lock *model.Lock) error

	// Delete deletes lock with given `resourceID`.
	//
	// If there is no such lock(by `lock`.ResourceID), then `ErrLockNotFound` is
	// returned.
	Delete(ctx context.Context, resourceID string) error

	// Get returns lock with given `resourceID`.
	//
	// If there is no such lock(by `lock`.ResourceID), then `ErrLockNotFound` is
	// returned.
	Get(ctx context.Context, resourceID string) (*model.Lock, error)

	// DeleteIfTokenMatches deletes lock with given `lock`.ResourceID only if
	// `lock`.Token fits.
	//
	// If there is no such `lock`(by `lock`.ResourceID), then
	// `ErrLockNotFound` is returned.
	//
	// If `lock`.Token doesn't fit, then `ErrWrongToken` is returned.
	DeleteIfTokenMatches(ctx context.Context, lock *model.Lock) error
}
