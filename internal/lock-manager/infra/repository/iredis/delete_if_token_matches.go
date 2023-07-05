package iredis

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/infra/repository"
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

func (s LockStorage) DeleteIfTokenMatches(ctx context.Context, resourceID, token string) error {
	script := redis.NewScript(deleteIfTokenMatchesScript)

	ec, err := script.Run(ctx, s.db,
		[]string{resourceID},
		[]string{token},
	).Result()
	if err != nil {
		s.l.Error(
			err,
			"resourceID", resourceID,
			"token", token,
		)

		return err
	}

	exitCode, ok := ec.(int64)
	if !ok {
		err = errors.New("type assertion error")
		s.l.Error(
			err,
			"resourceID", resourceID,
			"token", token,
			"ec", ec,
		)

		return err
	}
	switch exitCode {
	case InvalidTokenExitCode:
		err = repository.ErrInvalidToken
		s.l.Info(
			err.Error(),
			"resourceID", resourceID,
			"token", token,
		)
	case LockNotFoundExitCode:
		err = repository.ErrLockNotFound
		s.l.Info(
			err.Error(),
			"resourceID", resourceID,
			"token", token,
		)
	}

	return err
}
