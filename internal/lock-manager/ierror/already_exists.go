package ierror

import (
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

// NewAlreadyExists creates a new AlreadyExistsError object with the given message and enum value.
//
// Should be used to create a 'static' error, such as:
//
//	var ErrAlreadyExists = ierror.NewAlreadyExists("user with this login already exists", "USER_ALREADY_EXISTS")
func NewAlreadyExists(msg, enum string) AlreadyExistsErrorI {
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

// InstantiateAlreadyExists creates a new AlreadyExistsError from "err" with
// populated "duplicateID" field.
//
// Should be used when you want to create a 'dynamic' error in order to propagate
// ID of the duplicate via the error, such as:
//
//	if (...){
//		 return nil, ierror.InstantiateAlreadyExists(ErrAlreadyExists, id)
//	}
//
// Panics if "err" is nil.
func InstantiateAlreadyExists(
	err *AlreadyExistsError,
	dupID string,
) AlreadyExistsErrorI {
	if err == nil {
		panic("err cannot be nil")
	}

	return &AlreadyExistsError{
		duplicateID: dupID,
		APIError:    err.APIError,
	}
}

func (e AlreadyExistsError) DuplicateID() (string, bool) {
	return e.duplicateID, e.duplicateID != ""
}
