//go:build wireinject
// +build wireinject

package iprom

import (
	"github.com/google/wire"

	apputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app"
)

//nolint:gochecknoglobals // Wire Set
var (
	WireMetricSet = wire.NewSet(New, bind)
	bind          = wire.Bind(new(apputil.MetricI), new(*Metric))
)
