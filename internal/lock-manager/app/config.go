package app

import (
	redisconn "github.com/ruslanSorokin/lock-manager/internal/pkg/conn/redis"
	apputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app"
	grpcutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/grpc"
	promutil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/prom"
)

type Config struct {
	Redis      *redisconn.Config
	GRPC       *grpcutil.Config
	HTTPMetric *promutil.Config
	Ver        apputil.Ver
}
