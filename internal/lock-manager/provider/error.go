package provider

import (
	"errors"
	"fmt"
)

const errPrefix = "lock storage"

// Errors of LockProviderI.
var (
	ErrLockAlreadyExists = errors.New("lock already exists")
	ErrLockNotFound      = errors.New("lock is not found")
	ErrWrongToken        = errors.New("wrong token")
)

func Errf(err error) error {
	return fmt.Errorf("%s: %w", errPrefix, err)
}
