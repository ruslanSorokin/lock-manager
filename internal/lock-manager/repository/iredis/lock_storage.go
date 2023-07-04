package iredis

import (
	"github.com/go-logr/logr"
	"github.com/redis/go-redis/v9"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/repository"
)

// LockStorage stores logger & redis client.
type LockStorage struct {
	l  logr.Logger
	db *redis.Client
}

var _ repository.LockStorageI = (*LockStorage)(nil)

// NewLockStorage creates a new LockStorage.
func NewLockStorage(l logr.Logger, db *redis.Client) LockStorage {
	return LockStorage{
		l:  l,
		db: db,
	}
}
