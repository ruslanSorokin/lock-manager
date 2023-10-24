package ierror

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type InvalidArgumentError struct {
	APIError
}

var _ APIErrorI = (*InvalidArgumentError)(nil)

// NewInvalidArgument creates a new InvalidArgumentError with corresponding HTTP
// & GRPC status codes.
func NewInvalidArgument(msg, enum string) *InvalidArgumentError {
	return &InvalidArgumentError{
		APIError: APIError{
			msg:  msg,
			grpc: codes.InvalidArgument,
			http: http.StatusBadRequest,
			enum: enum,
		},
	}
}
