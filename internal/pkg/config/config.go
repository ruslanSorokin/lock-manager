package config

import (
	"path/filepath"

	"github.com/go-logr/logr"
	"github.com/spf13/viper"
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

// MustLoad either returns parsed config of type *T or panics.
func MustLoad[T any](log logr.Logger, app string, config configType) *T {
	path := filepath.Join(folder, app)
	viper := viper.New()
	var cfg T

	viper.SetConfigType(extension)
	viper.SetConfigName(string(config))
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Error(
			err,
			"unable to read config",
			"appName", app,
			"cfgName", config,
		)

		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Error(
			err,
			"unable to parse config",
			"appName", app,
			"cfgName", config,
		)

		panic(err)
	}

	return &cfg
}
