package service

import (
	"fmt"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/ierror"
)

const errPrefix = "lock service"

func Errf(err error) error {
	return fmt.Errorf("%s: %w", errPrefix, err)
}

var (
	ErrInvalidResourceID = ierror.NewInvalidArgument(
		"invalid resource ID",
		"INVALID_RESOURCE_ID",
	)
	ErrInvalidToken = ierror.NewInvalidArgument(
		"invalid token",
		"INVALID_TOKEN",
	)
)
