//go:build wireinject
// +build wireinject

package redisconn

import (
	"github.com/google/wire"
)

var (
	WireSet       = wire.NewSet(NewConn)
	WireConfigSet = wire.NewSet(NewConnFromConfig)
)
