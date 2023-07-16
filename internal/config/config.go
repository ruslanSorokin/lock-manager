package config

import (
	"path/filepath"

	"github.com/go-logr/logr"
	"github.com/ilyakaznacheev/cleanenv"
)

type configType string

const (
	folder    = "configs"
	extension = "yaml"
)

// Config type.
const (
	Local configType = "local"
	Dev   configType = "dev"
	Prod  configType = "prod"
)

const (
	configExtension = ".yaml"
)

const (
	logPathKey = "path"
)

// MustLoad either returns parsed config of type *T or panics.
func MustLoad[T any](log logr.Logger, config configType) *T {
	path := filepath.Join(folder, string(config)) + configExtension
	var cfg T

	err := cleanenv.ReadConfig(path, &cfg)
	log.Info(
		"trying to read config",
		logPathKey, path,
	)
	if err != nil {
		log.Error(
			err, "unable to read config",
			logPathKey, path,
		)
		panic(err)
	}

	return &cfg
}
