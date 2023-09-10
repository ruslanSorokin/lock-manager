package iprom

import (
	"github.com/google/wire"

	grpcutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/grpc"
)

//nolint:gochecknoglobals // Wire Set
var (
	WireRecoveryMetricSet = wire.NewSet(NewRecoveryMetric, bind)

	bind = wire.Bind(new(grpcutil.RecoveryMetricI), new(*RecoveryMetric))
)
