package ierror

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type invalidArgumentError struct {
	*logicalError
}

var _ errorI = (*invalidArgumentError)(nil)

// NewInvalidArgument creates a new invalidArgumentError with corresponding HTTP
// and GRPC status codes.
func NewInvalidArgument(msg, apiStCode string) error {
	return &invalidArgumentError{
		logicalError: &logicalError{
			msg:  msg,
			grpc: codes.InvalidArgument,
			http: http.StatusBadRequest,
			api:  apiStCode,
		},
	}
}
