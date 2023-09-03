package iredis

import (
	"context"
	"errors"

	redis "github.com/redis/go-redis/v9"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/model"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
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

func (s LockStorage) DeleteIfTokenMatches(ctx context.Context, lock *model.Lock) error {
	script := redis.NewScript(deleteIfTokenMatchesScript)

	ec, err := script.Run(ctx, s.conn.DB,
		[]string{lock.ResourceID()},
		lock.Token(),
	).Result()
	if err != nil {
		s.l.Error(err,
			"resourceID", lock.ResourceID,
			"token", lock.Token,
		)
		return provider.Errf(err)
	}

	exitCode, ok := ec.(int64)
	if !ok {
		err = errors.New("type assertion error")
		s.l.Error(err,
			"resourceID", lock.ResourceID,
			"token", lock.Token,
			"ec", ec,
		)
		return provider.Errf(err)
	}
	switch exitCode {
	case InvalidTokenExitCode:
		err = provider.ErrWrongToken
		return provider.Errf(err)

	case LockNotFoundExitCode:
		err = provider.ErrLockNotFound
		return provider.Errf(err)

	default:
		return nil
	}
}
