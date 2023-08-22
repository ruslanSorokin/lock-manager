//go:build wireinject
// +build wireinject

package service

import "github.com/google/wire"

//nolint:gochecknoglobals // Wire Sets
var (
	Set           = wire.NewSet(New, bind)
	FromConfigSet = wire.NewSet(NewFromConfig, bind)
	DefaultSet    = wire.NewSet(Default, bind)

	bind = wire.Bind(new(LockServiceI), new(*LockService))
)
