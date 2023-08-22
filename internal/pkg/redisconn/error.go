package redisconn

import (
	"errors"
	"fmt"
)

var ErrUnableToPing = errors.New("unable to ping redis")

func wrapErr(err error) error {
	return fmt.Errorf("%s: %w", errPrefix, err)
}
