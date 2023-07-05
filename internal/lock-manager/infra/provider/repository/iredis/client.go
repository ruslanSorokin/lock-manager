package iredis

import (
	"context"
	"time"

	"github.com/go-logr/logr"
	"github.com/redis/go-redis/v9"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/provider"
)

// Config stores all configurable parameters for redis connection.
type Config struct {
	URI      string
	Username string
	Password string
	DB       uint
}

const connectionTimeout = 10 * time.Second

// NewClient creates a new instance of redis client.
func NewClient(l logr.Logger, uri, uname, pword string, db uint) (*redis.Client, error) {
	l.Info(
		"trying to connect to redis",
		"uri", uri,
		"db", db,
	)

	cl := redis.NewClient(
		&redis.Options{
			Addr:     uri,
			Username: uname,
			Password: pword,
			DB:       int(db),
		},
	)

	ctx, cancel := context.WithTimeout(context.Background(), connectionTimeout)
	defer cancel()

	err := cl.Ping(ctx).Err()
	if err != nil {
		l.Error(
			err,
			"unable to ping redis",
		)

		return nil, provider.Errf(err)
	}

	l.Info(
		"connected to redis",
	)

	return cl, nil
}

// NewClientFromConfig creates a new instance of redis client from config.
func NewClientFromConfig(l logr.Logger, cfg Config) (*redis.Client, error) {
	return NewClient(l, cfg.URI, cfg.Username, cfg.Password, cfg.DB)
}
