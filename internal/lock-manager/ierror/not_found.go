package ierror

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type notFoundError struct {
	*logicalError
}

var _ errorI = (*notFoundError)(nil)

// NewNotFound creates a new notFoundError with corresponding HTTP
// and GRPC status codes.
func NewNotFound(msg, apiStCode string) error {
	return &notFoundError{
		logicalError: &logicalError{
			msg:  msg,
			grpc: codes.NotFound,
			http: http.StatusNotFound,
			api:  apiStCode,
		},
	}
}
