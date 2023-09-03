package iprom

import (
	"github.com/google/wire"

	"github.com/ruslanSorokin/lock-manager/internal/pkg/grpcutil"
)

//nolint:gochecknoglobals // Wire Set
var (
	WireRecoveryMetricSet = wire.NewSet(NewRecoveryMetric, bind)

	bind = wire.Bind(new(grpcutil.RecoveryMetricI), new(*RecoveryMetric))
)
