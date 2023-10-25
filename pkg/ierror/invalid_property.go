package ierror

import (
	"errors"
	"fmt"
)

type InvalidPropertyErrorI interface {
	APIErrorI

	Violation() (string, bool)
}

type InvalidPropertyError struct {
	PropertyError

	violation string
}

var _ InvalidPropertyErrorI = (*InvalidPropertyError)(nil)

// NewInvalidProperty creates a new InvalidPropertyError with corresponding status codes:
//
// - HTTP: 409
//
// - GRPC: 6 .
func NewInvalidProperty(prop, enum string) *InvalidPropertyError {
	msg := fmt.Sprintf("invalid property: %s", prop)
	return &InvalidPropertyError{
		violation: "",
		PropertyError: PropertyError{
			InvalidArgumentError: *NewInvalidArgument(msg, enum),
			property:             prop,
		},
	}
}

// InstantiateInvalidProperty creates a new InvalidPropertyError from "err" with
// populated "violation" field.
//
// Panics if "err" is nil.
func InstantiateInvalidProperty(
	err *InvalidPropertyError,
	vio string,
) *InvalidPropertyError {
	if err == nil {
		panic("err cannot be nil")
	}

	return &InvalidPropertyError{
		violation:     vio,
		PropertyError: err.PropertyError,
	}
}

// Violation returns property's violation.
func (e InvalidPropertyError) Violation() (string, bool) {
	return e.violation, e.violation != ""
}

func (e InvalidPropertyError) Is(target error) bool {
	t := NewInvalidProperty("", "")
	return errors.As(target, &t) && t.APIError == e.APIError
}
