//go:build wireinject
// +build wireinject

package igrpc

import "github.com/google/wire"

//nolint:gochecknoglobals // Wire Set
var (
	WireLockHandlerSet = wire.NewSet(NewLockHandler, bind)
	bind               = wire.Bind(new(LockHandlerI), new(LockHandler))
)
