//go:build wireinject
// +build wireinject

package iredis

import (
	"github.com/google/wire"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
)

//nolint:gochecknoglobals // Wire Set
var (
	WireLockStorageSet = wire.NewSet(NewLockStorage, bind)
	bind               = wire.Bind(new(provider.LockProviderI), new(*LockStorage))
)
