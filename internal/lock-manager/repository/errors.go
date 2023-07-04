package repository

import "errors"

var (
	ErrLockAlreadyExists = errors.New("LockStorage: lock already exists")
	ErrLockNotFound      = errors.New("LockStorage: lock is not found")
	ErrInvalidToken      = errors.New("LockStorage: invalid token")
)
