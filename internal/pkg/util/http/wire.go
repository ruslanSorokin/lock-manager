//go:build wireinject
// +build wireinject

package httputil

import (
	"github.com/google/wire"
)

//nolint:gochecknoglobals // Wire Sets
var (
	WireHandlerSet = wire.NewSet(NewHandler, bind)

	bind = wire.Bind(new(HandlerI), new(*Handler))
)
