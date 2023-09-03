//go:build wireinject
// +build wireinject

package iprom

import (
	"github.com/google/wire"

	"github.com/ruslanSorokin/lock-manager/internal/pkg/apputil"
)

//nolint:gochecknoglobals // Wire Set
var (
	WireMetricSet = wire.NewSet(New, bind)
	bind          = wire.Bind(new(apputil.MetricI), new(*Metric))
)
