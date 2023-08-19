package main

import (
	"github.com/ruslanSorokin/lock-manager/internal/infra/handler/igrpc"
	"github.com/ruslanSorokin/lock-manager/internal/service"
	"github.com/ruslanSorokin/lock-manager/pkg/promutil"
	"github.com/ruslanSorokin/lock-manager/pkg/redisutil"
)

type Config struct {
	Repository struct {
		Redis redisutil.Config `yaml:"redis"`
	} `yaml:"repository"`
	Handler struct {
		GRPC igrpc.Config `yaml:"grpc"`
	} `yaml:"handler"`
	Service       service.Config `yaml:"service"`
	Observability struct {
		Pull struct {
			Metric promutil.Config `yaml:"metric"`
		} `yaml:"pull"`
	} `yaml:"observability"`
	App struct {
		Environment string `yaml:"environment"`
		Version     string `yaml:"version"`
	} `yaml:"application"`
}
