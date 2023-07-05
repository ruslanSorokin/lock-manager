package main

import "github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/provider/repository/iredis"

type Config struct {
	redis iredis.Config
}
