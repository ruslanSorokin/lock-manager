package service

import (
	"errors"
	"fmt"
)

const errPrefix = "lock service"

func Errorf(err error) error {
	return fmt.Errorf("%s: %w", errPrefix, err)
}

var (
	ErrInvalidResourceID = errors.New("invalid resource ID")
	ErrInvalidToken      = errors.New("invalid token")
)
