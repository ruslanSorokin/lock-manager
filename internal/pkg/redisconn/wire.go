//go:build wireinject
// +build wireinject

package redisconn

import (
	"github.com/google/wire"
)

var (
	Set       = wire.NewSet(NewConn)
	ConfigSet = wire.NewSet(NewConnFromConfig)
)
