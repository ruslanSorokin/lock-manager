package service

import (
	"fmt"

	"github.com/ruslanSorokin/lock-manager/pkg/ierror"
)

func Errf(err error) error {
	return fmt.Errorf("lock service: %w", err)
}

var (
	ErrInvalidResourceID = ierror.NewInvalidArgument(
		"invalid resource ID",
		"INVALID_RESOURCE_ID")

	ErrInvalidToken = ierror.NewInvalidArgument(
		"invalid token",
		"INVALID_TOKEN")
)
