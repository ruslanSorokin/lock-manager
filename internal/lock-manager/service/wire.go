//go:build wireinject
// +build wireinject

package service

import "github.com/google/wire"

//nolint:gochecknoglobals // Wire Sets
var (
	WireSet           = wire.NewSet(New, bind)
	WireFromConfigSet = wire.NewSet(NewFromConfig, bind)
	WireDefaultSet    = wire.NewSet(Default, bind)

	bind = wire.Bind(new(LockServiceI), new(*LockService))
)
