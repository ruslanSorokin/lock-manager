//go:build wireinject
// +build wireinject

package iprom

import (
	"github.com/google/wire"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/grpcutil"
)

//nolint:gochecknoglobals // Wire Set
var (
	RecoveryMetricSet = wire.NewSet(NewRecoveryMetric, bind)

	bind = wire.Bind(new(grpcutil.RecoveryMetricI), new(*RecoveryMetric))
)
