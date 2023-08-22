package apputil

import (
	"github.com/pkg/errors"
)

type (
	Env string
	Ver string
)

// ErrInvalidEnv returns if given string cannot be matched to env type.
var ErrInvalidEnv = errors.New("invalid env")

// MustParseEnv parses env and trying to match it to one of the Env type.
//
// Panics if env cannot be matched.
func MustParseEnv(env string) Env {
	e, err := ParseEnv(env)
	if err != nil {
		panic(err)
	}
	return e
}

// ParseEnv parses env and trying to match it to one of the Env type.
//
// Returns ErrInvalidEnv if env cannot be matched.
func ParseEnv(env string) (Env, error) {
	switch env {
	case "local", "loc", "lcl", "l":
		return Local, nil
	case "development", "dev", "dvl", "d":
		return Dev, nil
	case "production", "prod", "prd", "p":
		return Prod, nil
	default:
		return "", errors.Wrap(ErrInvalidEnv, env)
	}
}

// Environment type.
const (
	Local Env = "local"
	Dev   Env = "development"
	Prod  Env = "production"
)
