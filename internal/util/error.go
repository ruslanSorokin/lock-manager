package util

import "fmt"

func NewErrWrapper(prefix string) func(error) error {
	return func(e error) error {
		if e != nil {
			return fmt.Errorf("%s: %w", prefix, e)
		}

		return nil
	}
}
