//go:build wireinject
// +build wireinject

package iprom

import (
	"github.com/google/wire"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/imetric"
)

//nolint:gochecknoglobals // Wire Set
var (
	Set  = wire.NewSet(New, bind)
	bind = wire.Bind(new(imetric.ServiceMetricI), new(*Metric))
)
