//go:build wireinject
// +build wireinject

package ifiber

import "github.com/google/wire"

//nolint:gochecknoglobals // Wire Set
var (
	WireLockHandlerSet = wire.NewSet(NewLockHandler)
)
