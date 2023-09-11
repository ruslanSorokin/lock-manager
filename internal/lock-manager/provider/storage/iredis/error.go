package iredis

import (
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
)

var (
	ErrLockAlreadyExists = provider.ErrLockAlreadyExists
	ErrLockNotFound      = provider.ErrLockNotFound
	ErrWrongToken        = provider.ErrWrongToken
)

func Errf(err error) error {
	return provider.Errf(err)
}
