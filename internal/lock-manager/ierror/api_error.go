package ierror

import (
	"google.golang.org/grpc/codes"
)

type apiErrorI interface {
	error

	GRPCStatusCode() codes.Code
	HTTPStatusCode() int
	APIStatusCode() string
}

type apiError struct {
	msg  string
	grpc codes.Code
	http int
	api  string
}

var _ apiErrorI = (*apiError)(nil)

// New creates new a apiError that meets the apiErrorI interface.
func New(msg string, grpc codes.Code, http int, api string) error {
	return &apiError{msg: msg, grpc: grpc, http: http, api: api}
}

func (e *apiError) Error() string { return e.msg }

func (e *apiError) GRPCStatusCode() codes.Code { return e.grpc }

func (e *apiError) HTTPStatusCode() int { return e.http }

func (e *apiError) APIStatusCode() string { return e.api }
