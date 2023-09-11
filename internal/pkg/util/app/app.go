package apputil

import (
	"errors"
	"fmt"
)

type (
	Name string
	Env  string
	Ver  string
)

func (n Name) String() string {
	return string(n)
}

func (e Env) String() string {
	return string(e)
}

func (v Ver) String() string {
	return string(v)
}

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
	case "development", "dev", "d":
		return Dev, nil
	case "testing", "test", "t":
		return Test, nil
	case "production", "prod", "p":
		return Prod, nil
	default:
		return "", fmt.Errorf("%w: %s", ErrInvalidEnv, env)
	}
}

// Environment type.
const (
	Dev  Env = "development"
	Test Env = "test"
	Prod Env = "production"
)
