package main

import (
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/app"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/apputil"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/grpcutil"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/promutil"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/redisconn"
)

type Config struct {
	Repository struct {
		Redis *redisconn.Config `yaml:"redis"`
	} `yaml:"repository"`
	Handler struct {
		GRPC *grpcutil.Config `yaml:"grpc"`
	} `yaml:"handler"`
	Service       *service.Config `yaml:"service"`
	Observability struct {
		Pull struct {
			Metric *promutil.Config `yaml:"metric"`
		} `yaml:"pull"`
	} `yaml:"observability"`
	App struct {
		Environment string `yaml:"environment"`
		Version     string `yaml:"version"`
	} `yaml:"application"`
}

func (c Config) ToAppConfig() *app.Config {
	return &app.Config{
		Redis:       c.Repository.Redis,
		GRPC:        c.Handler.GRPC,
		LockService: c.Service,
		HTTPMetric:  c.Observability.Pull.Metric,
		Ver:         apputil.Ver(c.App.Version),
	}
}
