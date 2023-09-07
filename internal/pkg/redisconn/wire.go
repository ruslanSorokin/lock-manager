//go:build wireinject
// +build wireinject

package redisconn

import (
	"github.com/google/wire"
)

var (
	WireSet       = wire.NewSet(New)
	WireConfigSet = wire.NewSet(NewFromConfig)
)
