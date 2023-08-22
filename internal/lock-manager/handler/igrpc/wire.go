//go:build wireinject
// +build wireinject

package igrpc

import "github.com/google/wire"

//nolint:gochecknoglobals // Wire Set
var (
	LockHandlerSet = wire.NewSet(NewLockHandler, bind)
	bind           = wire.Bind(new(LockHandlerI), new(LockHandler))
)
