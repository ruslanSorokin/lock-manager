//go:build wireinject
// +build wireinject

package service

import "github.com/google/wire"

//nolint:gochecknoglobals // Wire Set
var (
	WireSet = wire.NewSet(New, bind)

	bind = wire.Bind(new(LockServiceI), new(*LockService))
)
