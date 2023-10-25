package ierror

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type NotFoundError struct {
	APIError
}

var _ APIErrorI = (*NotFoundError)(nil)

// NewNotFound creates a new NotFoundError with corresponding status codes:
//
// - HTTP: 404
//
// - GRPC: 5 .
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
