package main

import "github.com/ruslanSorokin/lock-manager/internal/infra/provider/repository/iredis"

type Config struct {
	redis iredis.Config
}
