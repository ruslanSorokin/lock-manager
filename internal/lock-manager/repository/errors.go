package repository

import "errors"

var (
	ErrLockAlreadyExists = errors.New("lock storage: lock already exists")
	ErrLockNotFound      = errors.New("lock storage: lock is not found")
	ErrInvalidToken      = errors.New("lock storage: invalid token")
)
