package util

import (
	"fmt"

	"github.com/ory/dockertest"
)

func Must[T any](t T, u error) T {
	if u != nil {
		panic(u)
	}
	return t
}

func NewPool() (*dockertest.Pool, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		err = fmt.Errorf("%s: %w", "could not construct the pool", err)
		return nil, err
	}

	err = pool.Client.Ping()
	if err != nil {
		err = fmt.Errorf("%s: %w", "could not connect to Docker", err)
		return nil, err
	}
	return pool, nil
}
