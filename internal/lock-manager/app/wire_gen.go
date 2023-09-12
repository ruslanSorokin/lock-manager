// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/go-logr/logr"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/handler/igrpc"
	iprom2 "github.com/ruslanSorokin/lock-manager/internal/lock-manager/metric/iprom"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider/storage/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/conn/redis"
	iprom3 "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app/iprom"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/util/grpc"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/util/grpc/iprom"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/util/http"
	"net/http"
)

// Injectors from wire.go:

func Wire(logger logr.Logger, config *Config) (*App, func(), error) {
	appWireConfig, err := toWireConfig(config)
	if err != nil {
		return nil, nil, err
	}
	redisconnConfig := appWireConfig.Redis
	conn, cleanup, err := redisconn.WireProvide(logger, redisconnConfig)
	if err != nil {
		return nil, nil, err
	}
	registry := prometheus.NewRegistry()
	serverMetrics := grpcutil.NewProcessingTimeHistogram(registry)
	serveMux := http.NewServeMux()
	server := provideHTTPServer()
	loggingLogger := grpcutil.NewInterceptorLogger(logger)
	recoveryMetric := iprom.NewRecoveryMetric(registry)
	v := grpcutil.NewPanicRecoveryHandler(logger, recoveryMetric)
	v2 := grpcutil.WireProvideInterceptors(loggingLogger, v, serverMetrics)
	grpcutilConfig := appWireConfig.GRPC
	grpcServer := grpcutil.WireProvideServer(v2, grpcutilConfig)
	validate := validator.New()
	lockStorage := iredis.NewLockStorage(logger, conn)
	ipromMetric := iprom2.New(registry)
	lockService := service.New(logger, validate, lockStorage, ipromMetric)
	grpcutilHandler := grpcutil.NewHandlerFromConfig(grpcServer, logger, grpcutilConfig)
	lockHandler := igrpc.NewLockHandler(grpcutilHandler, logger, lockService)
	httputilConfig := appWireConfig.Pull
	httputilHandler := httputil.NewHandlerFromConfig(logger, server, serveMux, httputilConfig)
	metric2 := iprom3.New(registry)
	env := appWireConfig.Environment
	ver := appWireConfig.Version
	appApp := New(config, logger, conn, registry, serverMetrics, serveMux, server, grpcServer, lockService, lockHandler, httputilHandler, metric2, env, ver)
	return appApp, func() {
		cleanup()
	}, nil
}

// wire.go:

func provideHTTPServer() *http.Server {
	return &http.Server{}
}
