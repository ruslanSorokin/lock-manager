package grpcutil

import (
	promgrpc "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func WireProvideDefaultServerOpts(
	log logging.Logger,
	recoveryHandler func(any) error,
	metric *promgrpc.ServerMetrics,
) []grpc.ServerOption {
	unaryInters := []grpc.UnaryServerInterceptor{
		metric.UnaryServerInterceptor(),
		logging.UnaryServerInterceptor(log),
		recovery.UnaryServerInterceptor(
			recovery.WithRecoveryHandler(recoveryHandler),
		),
	}

	streamInters := []grpc.StreamServerInterceptor{
		metric.StreamServerInterceptor(),
		logging.StreamServerInterceptor(log),
		recovery.StreamServerInterceptor(
			recovery.WithRecoveryHandler(recoveryHandler),
		),
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(unaryInters...),
		grpc.ChainStreamInterceptor(streamInters...),
	}

	return opts
}

func WireProvideServer(
	opts []grpc.ServerOption,
	cfg *Config,
) *grpc.Server {
	opts = append(opts,
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     cfg.Conn.MaxIdle,
			MaxConnectionAge:      cfg.Conn.MaxAge,
			MaxConnectionAgeGrace: cfg.Conn.Grace,

			Time:    cfg.Ping.After,
			Timeout: cfg.Ping.Timeout,
		}))
	return grpc.NewServer(opts...)
}
