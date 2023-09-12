//go:build wireinject
// +build wireinject

package httputil

import (
	"github.com/google/wire"
)

var (
	WireHandlerSet           = wire.NewSet(NewHandler)
	WireHandlerFromConfigSet = wire.NewSet(NewHandlerFromConfig)
)
