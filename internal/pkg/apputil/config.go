package apputil

import (
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	folder = "configs"
)

const (
	configExtension = ".yaml"
)

// Load returns parsed config of type *T or error.
func Load[T any](env Env) (*T, error) {
	path := filepath.Join(folder, string(env)) + configExtension
	var cfg T

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// MustLoad either returns parsed config of type *T or panics.
func MustLoad[T any](env Env) *T {
	cfg, err := Load[T](env)
	if err != nil {
		panic(err)
	}
	return cfg
}
