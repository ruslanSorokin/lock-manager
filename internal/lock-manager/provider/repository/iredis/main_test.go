package iredis_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/ory/dockertest"
	"github.com/rs/zerolog"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider/repository/iredis"
	"github.com/ruslanSorokin/lock-manager/internal/pkg/redisconn"
)

//nolint:gochecknoglobals // Using global var in tests
var redisCl *redisconn.Conn

//nolint:gochecknoglobals // Using global var in tests
var log logr.Logger

const (
	redisImageName    = "redis"
	redisImageVersion = "7.0.10"
)

const (
	redisIP   = "localhost"
	redisPort = "6379/tcp"
)

func flushStorage(conn *redisconn.Conn) error {
	return conn.DB.FlushAll(context.Background()).Err()
}

func TestMain(m *testing.M) {
	zl := zerolog.New(os.Stdout)
	zl = zl.With().Logger()

	log = zerologr.New(&zl).V(3)

	flag.Parse()
	if testing.Short() {
		os.Exit(m.Run())
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Error(err, "could not construct the pool")
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Error(err, "could not connect to Docker")
	}

	resource, err := pool.Run(
		redisImageName, redisImageVersion, nil,
	)
	if err != nil {
		log.Error(err, "could not start the resource")
	}

	cfg := &redisconn.Config{
		URI:      fmt.Sprintf("%s:%s", redisIP, resource.GetPort(redisPort)),
		Username: "",
		Password: "",
		DB:       0,
	}

	if err = pool.Retry(func() error {
		redisCl, err = redisconn.NewConnFromConfig(cfg)
		redisDB := redisCl.DB
		if err != nil {
			return err
		}

		return redisDB.Ping(context.TODO()).Err()
	}); err != nil {
		log.Error(err, "could not connect to Docker")
	}

	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Error(err, "could not purge the resource")
	}

	os.Exit(code)
}

func TestIntegrationRedisLockStorage(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	lockStorageGetter := func() func() provider.LockProviderI {
		return func() provider.LockProviderI {
			return iredis.NewLockStorage(log, redisCl)
		}
	}()

	dbFlusherGetter := func() func() error {
		return func() error {
			return flushStorage(redisCl)
		}
	}()

	provider.RunLockStorageTests(
		t,
		lockStorageGetter,
		dbFlusherGetter,
	)
}
