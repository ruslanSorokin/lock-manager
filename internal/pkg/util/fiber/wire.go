//go:build wireinject
// +build wireinject

package fiberutil

import "github.com/google/wire"

var (
	WireHandlerFromConfigSet = wire.NewSet(NewHandlerFromConfig, bind)

	bind = wire.Bind(new(HandlerI), new(*Handler))
)
