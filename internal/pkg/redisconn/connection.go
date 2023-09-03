package redisconn

import (
	"context"
	"errors"
	"time"

	redis "github.com/redis/go-redis/v9"
)

const (
	// TODO: move timeout and retry options to config
	_connTimeout = 10 * time.Second

	errPrefix = "redisutil"
)

type Conn struct {
	DB *redis.Client
}

// NewConn creates a new redis Conn.
func NewConn(uri, uname, pword string, db uint) (*Conn, error) {
	c := &Conn{
		DB: redis.NewClient(
			&redis.Options{
				Addr:     uri,
				Username: uname,
				Password: pword,
				DB:       int(db),
			},
		),
	}

	ctx, cancel := context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()
	err := c.DB.Ping(ctx).Err()
	if err != nil {
		err := errors.Join(err, ErrUnableToPing)
		return nil, wrapErr(err)
	}

	return c, nil
}

// NewConnFromConfig creates a new instance of redis client from config.
func NewConnFromConfig(cfg *Config) (*Conn, error) {
	return NewConn(cfg.URI, cfg.Username, cfg.Password, cfg.DB)
}

func (c Conn) HealthCheck() bool {
	ctx, cancel := context.WithTimeout(context.Background(), _connTimeout)
	defer cancel()
	err := c.DB.Ping(ctx).Err()

	return err == nil
}

func (c Conn) Close() error {
	return c.DB.Close()
}
