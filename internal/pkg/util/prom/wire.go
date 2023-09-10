//go:build wireinject
// +build wireinject

package promutil

import (
	"github.com/google/wire"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	WireHandlerSet           = wire.NewSet(NewHandler)
	WireHandlerFromConfigSet = wire.NewSet(NewHandlerFromConfig)
	WireRegistrySet          = wire.NewSet(prometheus.NewRegistry, bind)

	bind = wire.Bind(new(prometheus.Registerer), new(*prometheus.Registry))
)
