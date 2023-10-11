package iredis_test

import (
	"context"
	"flag"
	"fmt"
	"testing"

	"github.com/go-logr/logr"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/suite"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider/storage/iredis"
	providertest "github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider/test"
	redisconn "github.com/ruslanSorokin/lock-manager/internal/pkg/conn/redis"
	util "github.com/ruslanSorokin/lock-manager/internal/pkg/util"
)

const (
	redisImageName    = "redis"
	redisImageVersion = "7.0.10"
)

const (
	redisIP   = "localhost"
	redisPort = "6379/tcp"
)

type IntegrationSuite struct {
	*suite.Suite
	*providertest.PSuite

	resource *dockertest.Resource
	pool     *dockertest.Pool
	conn     *redisconn.Conn
}

func (s *IntegrationSuite) SetupSuite() {
	t := s.T()

	flag.Parse()
	if testing.Short() {
		t.Skip()
	}

	p, err := util.NewPool()
	if err != nil {
		t.Error(err)
	}
	s.pool = p

	r, err := LaunchRedisContainer(p, redisImageVersion)
	if err != nil {
		t.Error(err)
	}
	s.resource = r

	cfg := &redisconn.Config{
		URI:      fmt.Sprintf("%s:%s", redisIP, r.GetPort(redisPort)),
		Username: "",
		Password: "",
		DB:       0,
	}

	c, err := redisconn.New(cfg)
	if err != nil {
		t.Error(err)
	}
	s.conn = c

	ls := iredis.NewLockStorage(logr.Discard(), c)
	s.PSuite = providertest.NewSuite(s, ls)

	s.PSuite.Provider = ls
}

func (s *IntegrationSuite) TearDownSuite() {
	t := s.T()

	err := RemoveRedisContainer(s.pool, s.resource)
	if err != nil {
		t.Error(err)
	}
}

func (s *IntegrationSuite) TearDownTest() {
	t := s.T()

	err := s.conn.DB.FlushAll(context.TODO()).Err()
	if err != nil {
		t.Error(err)
	}
}

func LaunchRedisContainer(
	p *dockertest.Pool,
	tag string,
) (*dockertest.Resource, error) {
	resource, err := p.Run(redisImageName, tag, nil)
	if err != nil {
		err = fmt.Errorf("%s: %w", "could not start the resource", err)
		return nil, err
	}
	return resource, nil
}

func RemoveRedisContainer(p *dockertest.Pool, r *dockertest.Resource) error {
	if err := p.Purge(r); err != nil {
		return fmt.Errorf("%s: %w", "could not purge the resource", err)
	}
	return nil
}

func TestIntegrationRedisLockStorage(t *testing.T) {
	suite.Run(t, &IntegrationSuite{
		Suite: &suite.Suite{},
	})
}
