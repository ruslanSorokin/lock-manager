package ierror

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"
)

type AlreadyExistsErrorI interface {
	APIErrorI

	DuplicateID() (string, bool)
}

type AlreadyExistsError struct {
	APIError

	duplicateID string
}

var _ AlreadyExistsErrorI = (*AlreadyExistsError)(nil)

// NewInvalidArgument creates a new InvalidArgumentError with corresponding status codes:
//
// - HTTP: 409
//
// - GRPC: 6 .
func NewAlreadyExists(msg, enum string) *AlreadyExistsError {
	return &AlreadyExistsError{
		duplicateID: "",
		APIError: APIError{
			msg:  msg,
			grpc: codes.AlreadyExists,
			http: http.StatusConflict,
			enum: enum,
		},
	}
}

// InstantiateAlreadyExists creates a copy of AlreadyExistsError based on "err" with
// populated "duplicateID" field.
//
// Panics if "err" is nil.
func InstantiateAlreadyExists(
	err *AlreadyExistsError,
	dupID string,
) *AlreadyExistsError {
	if err == nil {
		panic("err cannot be nil")
	}

	return &AlreadyExistsError{
		duplicateID: dupID,
		APIError:    err.APIError,
	}
}

// DuplicateID returns duplicate ID if field is populated.
func (e AlreadyExistsError) DuplicateID() (string, bool) {
	return e.duplicateID, e.duplicateID != ""
}

func (e AlreadyExistsError) Is(target error) bool {
	t := NewAlreadyExists("", "")
	return errors.As(target, &t) && t.APIError == e.APIError
}
