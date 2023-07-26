package config

import (
	"path/filepath"

	"github.com/go-logr/logr"
	"github.com/ilyakaznacheev/cleanenv"
)

type Type string

const (
	folder = "configs"
)

// Config type.
const (
	Local Type = "local"
	Dev   Type = "dev"
	Prod  Type = "prod"
)

const (
	configExtension = ".yaml"
)

const (
	logKeyPath       = "path"
	logKeyConfigType = "config type"
)

// MustLoad either returns parsed config of type *T or panics.
func MustLoad[T any](log logr.Logger, config Type) *T {
	path := filepath.Join(folder, string(config)) + configExtension
	var cfg T

	err := cleanenv.ReadConfig(path, &cfg)
	log.Info(
		"trying to read config",
		logKeyConfigType, config,
		logKeyPath, path,
	)
	if err != nil {
		log.Error(
			err, "unable to read config",
			logKeyConfigType, config,
			logKeyPath, path,
		)
		panic(err)
	}

	log.Info(
		"read config",
		logKeyConfigType, config,
		logKeyPath, path,
	)
	return &cfg
}
