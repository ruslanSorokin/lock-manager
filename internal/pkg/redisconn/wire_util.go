package redisconn

import "github.com/go-logr/logr"

func WireProvide(log logr.Logger, cfg *Config) (*Conn, func(), error) {
	conn, err := NewConnFromConfig(cfg)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		if err := conn.Close(); err != nil {
			log.Error(err, "redis.Close() error")
		}
	}
	return conn, cleanup, nil
}
