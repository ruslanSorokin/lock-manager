package ierror

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

type invalidArgumentError struct {
	*apiError
}

var _ apiErrorI = (*invalidArgumentError)(nil)

// NewInvalidArgument creates a new invalidArgumentError with corresponding HTTP
// and GRPC status codes.
func NewInvalidArgument(msg, apiStCode string) error {
	return &invalidArgumentError{
		apiError: &apiError{
			msg:  msg,
			grpc: codes.InvalidArgument,
			http: http.StatusBadRequest,
			api:  apiStCode,
		},
	}
}
