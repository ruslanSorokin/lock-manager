package ierror

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type NotFoundError struct {
	APIError
}

var _ APIErrorI = (*NotFoundError)(nil)

// NewNotFound creates a new NotFoundError with corresponding HTTP
// & GRPC status codes.
func NewNotFound(msg, enum string) *NotFoundError {
	return &NotFoundError{
		APIError: APIError{
			msg:  msg,
			grpc: codes.NotFound,
			http: http.StatusNotFound,
			enum: enum,
		},
	}
}
