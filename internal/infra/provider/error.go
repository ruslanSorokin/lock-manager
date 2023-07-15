package provider

import (
	"errors"

	"github.com/ruslanSorokin/lock-manager/internal/util"
)

const errPrefix = "lock storage"

// Errors of LockProviderI.
var (
	ErrLockAlreadyExists = errors.New("lock already exists")
	ErrLockNotFound      = errors.New("lock is not found")
	ErrWrongToken        = errors.New("wrong token")
)

//nolint:gochecknoglobals // Error wrapper
var Errf = util.NewErrWrapper(errPrefix)
