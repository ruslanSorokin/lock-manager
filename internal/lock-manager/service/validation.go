package service

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

type (
	resourceIDValidator func(string) error
	tokenValidator      func(string) error
)

func newResourceIDValidator(v *validator.Validate) resourceIDValidator {
	return func(r string) error {
		if err := v.Var(r, "required"); err != nil {
			err = fmt.Errorf("%w: %w", ErrInvalidResourceID, err)
			return err
		}
		return nil
	}
}

func newTokenValidator(v *validator.Validate) tokenValidator {
	return func(t string) error {
		if err := v.Var(t, "required,uuid4"); err != nil {
			err = fmt.Errorf("%w: %w", ErrInvalidToken, err)
			return err
		}
		return nil
	}
}
