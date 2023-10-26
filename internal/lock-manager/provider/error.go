package provider

import (
	"fmt"

	"github.com/ruslanSorokin/lock-manager/internal/pkg/ierror"
)

const errPrefix = "lock storage"

// Errors of LockProviderI.
var (
	ErrLockAlreadyExists = ierror.NewAlreadyExists(
		"lock already exists",
		"LOCK_ALREADY_EXISTS",
	)
	ErrLockNotFound = ierror.NewNotFound(
		"lock is not found",
		"LOCK_NOT_FOUND",
	)
	ErrWrongToken = ierror.NewInvalidArgument(
		"wrong token",
		"TOKEN_DOES_NOT_FIT",
	)
)

func Errf(err error) error {
	return fmt.Errorf("%s: %w", errPrefix, err)
}
