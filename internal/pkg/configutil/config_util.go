package configutil

import (
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/ruslanSorokin/lock-manager/internal/pkg/apputil"
)

const (
	configFolder = "config"
)

type (
	File string
)

func (f File) String() string {
	return string(f)
}

type Config[T any] struct {
	appName   apputil.Name
	appConfig *T
}

func NewConfig[T any](n apputil.Name) *Config[T] {
	c := &Config[T]{
		appName:   n,
		appConfig: new(T),
	}
	return c
}

func (c *Config[T]) AppConfig() *T {
	return c.appConfig
}

func (c *Config[T]) AppName() apputil.Name {
	return c.appName
}

// Load loads config from file or returns error.
// File will be searched as following `config/{c.appName}/{fname}`.
func (c *Config[T]) Load(fname File) error {
	path := filepath.Join(configFolder, c.appName.String(), fname.String())

	err := cleanenv.ReadConfig(path, c.appConfig)
	if err != nil {
		return err
	}

	err = cleanenv.ReadEnv(c.appConfig)
	if err != nil {
		return err
	}

	return nil
}

// MustLoad either returns parsed from file config of type *T or panics.
func (c *Config[T]) MustLoad(fname File) {
	err := c.Load(fname)
	if err != nil {
		panic(err)
	}
}
