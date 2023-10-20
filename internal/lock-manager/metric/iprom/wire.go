//go:build wireinject
// +build wireinject

package iprom

import (
	"github.com/google/wire"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/metric"
)

//nolint:gochecknoglobals // Wire Set
var (
	WireSet = wire.NewSet(New, bind)

	bind = wire.Bind(new(metric.ServiceMetricI), new(*Metric))
)
