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

var (
	//nolint:gochecknoglobals // Error wrapper
	Errf = util.NewErrWrapper(errPrefix)
)
