package repository

import (
	"context"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

// LockStorageI provides CRUD+Custom operations on the lock model.
type LockStorageI interface {
	// Create creates a new lock.
	//
	// If lock with such resourceID already exists, then ErrLockAlreadyExists is returned.
	Create(ctx context.Context, l model.Lock) error

	// Delete deletes lock with given resourceID.
	//
	// If there is no such lock, then ErrLockNotFound is returned.
	Delete(ctx context.Context, resourceID string) error

	// Get returns lock with given resourceID.
	//
	// If there is no such lock, then ErrLockNotFound is returned.
	Get(ctx context.Context, resourceID string) (*model.Lock, error)

	// GetAll returns slice of locks.
	//
	// If there is no locks at all, then the nil slice is returned.
	GetAll(ctx context.Context) []model.Lock

	// DeleteIfSecretKeyMatches deletes lock with given resourceID only if the
	// secret key is the same as the one that already in the storage.
	//
	// If there is no such lock, then ErrLockNotFound is returned.
	//
	// If secret key is not the same as the one that already in the storage, then
	// ErrInvalidSecretKey is returned.
	DeleteIfSecretKeyMatches(ctx context.Context, resourceID, secretKey string) error
}
