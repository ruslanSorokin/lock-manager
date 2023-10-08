package ierror

import (
	"google.golang.org/grpc/codes"
)

type errorI interface {
	error

	GRPCStatusCode() codes.Code
	HTTPStatusCode() int
	APIStatusCode() string
}

type logicalError struct {
	msg  string
	grpc codes.Code
	http int
	api  string
}

var _ errorI = (*logicalError)(nil)

func New(msg string, grpc codes.Code, http int) error {
	return &logicalError{msg: msg, grpc: grpc, http: http}
}

func (e *logicalError) Error() string { return e.msg }

func (e *logicalError) GRPCStatusCode() codes.Code { return e.grpc }

func (e *logicalError) HTTPStatusCode() int { return e.http }

func (e *logicalError) APIStatusCode() string { return e.api }
