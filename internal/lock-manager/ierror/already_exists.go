package ierror

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type alreadyExistsError struct {
	*logicalError
}

var _ errorI = (*alreadyExistsError)(nil)

// NewAlreadyExists creates a new alreadyExistsError with corresponding HTTP
// and GRPC status codes.
func NewAlreadyExists(msg, apiStCode string) error {
	return &alreadyExistsError{
		logicalError: &logicalError{
			msg:  msg,
			grpc: codes.AlreadyExists,
			http: http.StatusConflict,
			api:  apiStCode,
		},
	}
}
