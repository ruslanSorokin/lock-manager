//go:build wireinject
// +build wireinject

package httputil

import (
	"github.com/google/wire"
)

var (
	WireHandlerSet           = wire.NewSet(NewHandler, bind)
	WireHandlerFromConfigSet = wire.NewSet(NewHandlerFromConfig, bind)

	bind = wire.Bind(new(HandlerI), new(*Handler))
)
