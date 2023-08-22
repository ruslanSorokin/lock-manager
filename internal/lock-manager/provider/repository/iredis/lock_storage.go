package iredis

import (
	"github.com/go-logr/logr"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/redisconn"
)

// LockStorage stores logger & redis client.
type LockStorage struct {
	l    logr.Logger
	conn *redisconn.Conn
}

var _ provider.LockProviderI = (*LockStorage)(nil)

// NewLockStorage creates a new LockStorage.
func NewLockStorage(l logr.Logger, conn *redisconn.Conn) *LockStorage {
	return &LockStorage{
		l:    l,
		conn: conn,
	}
}
