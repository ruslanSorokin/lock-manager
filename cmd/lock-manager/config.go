package main

import (
	"github.com/ruslanSorokin/lock-manager/internal/infra/handler/igrpc"
	"github.com/ruslanSorokin/lock-manager/internal/infra/provider/repository/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/service"
)

type Config struct {
	Repository struct {
		Redis iredis.Config `yaml:"redis"`
	} `yaml:"repository"`
	Handler struct {
		GRPC igrpc.Config `yaml:"grpc"`
	} `yaml:"handler"`
	Service service.Config `yaml:"service"`
}
