package repository

import "errors"

// Errors that the LockStorage can return.
var (
	ErrLockAlreadyExists = errors.New("LockStorage: lock already exists")
	ErrLockNotFound      = errors.New("LockStorage: lock is not found")
	ErrInvalidToken      = errors.New("LockStorage: invalid token")
)
