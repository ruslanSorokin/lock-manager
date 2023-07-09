package iredis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/provider"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
)

//nolint:gosec // no credentials
const deleteIfTokenMatchesScript = `
		local token = redis.call("get", KEYS[1])
		if token == false then
			return 0
		end
		if token == ARGV[1] then
			return redis.call("del", KEYS[1])
		end

		return -999
`

// TODO: make auto substitution of the exit code into the Lua script
const (
	InvalidTokenExitCode = -999
	LockNotFoundExitCode = 0
)

func (s LockStorage) DeleteIfTokenMatches(ctx context.Context, lock model.Lock) error {
	script := redis.NewScript(deleteIfTokenMatchesScript)

	ec, err := script.Run(ctx, s.db,
		[]string{lock.ResourceID},
		[]string{lock.Token},
	).Result()
	if err != nil {
		s.l.Error(
			err,
			"resourceID", lock.ResourceID,
			"token", lock.Token,
		)

		return provider.Errf(err)
	}

	exitCode, ok := ec.(int64)
	if !ok {
		err = errors.New("type assertion error")
		s.l.Error(
			err,
			"resourceID", lock.ResourceID,
			"token", lock.Token,
			"ec", ec,
		)

		return provider.Errf(err)
	}
	switch exitCode {
	case InvalidTokenExitCode:
		err = provider.ErrInvalidToken
		s.l.Info(
			err.Error(),
			"resourceID", lock.ResourceID,
			"token", lock.Token,
		)
	case LockNotFoundExitCode:
		err = provider.ErrLockNotFound
		s.l.Info(
			err.Error(),
			"resourceID", lock.ResourceID,
			"token", lock.Token,
		)
	}

	return provider.Errf(err)
}
