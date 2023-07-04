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
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/repository"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/repository/iredis"
)

//nolint:gochecknoglobals // Using global var in tests
var redisCl *redis.Client

//nolint:gochecknoglobals // Using global var in tests
var log logr.Logger

const redisImageName = "redis"
const redisImageVersion = "7.0.10"

const redisIP = "localhost"
const redisPort = "6379/tcp"

func flushStorage(db *redis.Client) error {
	return db.FlushAll(context.Background()).Err()
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

	cfg := iredis.Config{
		URI:      fmt.Sprintf("%s:%s", redisIP, resource.GetPort(redisPort)),
		Username: "",
		Password: "",
		DB:       0,
	}

	if err = pool.Retry(func() error {
		redisCl, err = iredis.NewClientFromConfig(log, cfg)
		if err != nil {
			return err
		}

		return redisCl.Ping(context.TODO()).Err()
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

	lockStorageGetter := func() func() repository.LockStorageI {
		return func() repository.LockStorageI {
			return iredis.NewLockStorage(log, redisCl)
		}
	}()

	dbFlusherGetter := func() func() error {
		return func() error {
			return flushStorage(redisCl)
		}
	}()

	repository.RunLockStorageTests(
		t,
		lockStorageGetter,
		dbFlusherGetter,
	)
}
