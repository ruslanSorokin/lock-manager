package service

import (
	"fmt"

	"github.com/pkg/errors"
)

const errPrefix = "lock service"

func Errf(err error) error {
	return fmt.Errorf("%s: %w", errPrefix, err)
}

var (
	ErrInvalidResourceID = errors.New("invalid resource ID")
	ErrInvalidToken      = errors.New("invalid token")
)