package service

import (
	"errors"

	"github.com/ruslanSorokin/lock-manager/internal/pkg/util"
)

const errPrefix = "lock service"

//nolint:gochecknoglobals // Error wrapper
var Errf = util.NewErrWrapper(errPrefix)

var (
	ErrInvalidResourceID = errors.New("invalid resource ID")
	ErrInvalidToken      = errors.New("invalid token")
)
