package provider

import (
	"errors"

	"github.com/ruslanSorokin/lock-manager/internal/pkg/util"
)

const errPrefix = "lock storage"

// Errors of LockProviderI.
var (
	ErrLockAlreadyExists = errors.New("lock already exists")
	ErrLockNotFound      = errors.New("lock is not found")
	ErrInvalidToken      = errors.New("invalid token")
)

//nolint:gochecknoglobals // Error wrapper
var Errf = util.NewErrWrapper(errPrefix)
