package app

import (
	"github.com/ruslanSorokin/lock-manager/internal/pkg/apputil"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/grpcutil"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/promutil"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/redisconn"
)

type Config struct {
	Redis      *redisconn.Config
	GRPC       *grpcutil.Config
	HTTPMetric *promutil.Config
	Ver        apputil.Ver
}
