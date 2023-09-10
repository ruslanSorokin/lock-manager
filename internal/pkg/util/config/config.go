package configutil

import (
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"

	apputil "github.com/ruslanSorokin/lock-manager/internal/pkg/util/app"
)

const (
	configFolder = "config"
)

type File string

func (f File) String() string {
	return string(f)
}

// Load loads all configurable params for T from file and then rewrites them with env
// variables if any.
// Works on top of cleanenv struct tags.
// File will be searched as `config/{appName}/{fileName}`.
func Load[T any](c *T, appName apputil.Name, fileName File) error {
	path := filepath.Join(configFolder, appName.String(), fileName.String())
	return cleanenv.ReadConfig(path, c)
}

func MustLoad[T any](c *T, appName apputil.Name, fileName File) {
	if err := Load(c, appName, fileName); err != nil {
		panic(err)
	}
}
