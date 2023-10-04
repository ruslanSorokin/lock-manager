package ierror

import (
	"net/http"

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

type notFoundError logicalError

var _ errorI = (*notFoundError)(nil)

// NewNotFound creates a new notFoundError with corresponding HTTP
// and GPRC status codes.
func NewNotFound(msg, apiStCode string) error {
	return &notFoundError{
		msg:  msg,
		grpc: codes.NotFound,
		http: http.StatusNotFound,
		api:  apiStCode,
	}
}

func (e *notFoundError) Error() string { return e.msg }

func (e *notFoundError) GRPCStatusCode() codes.Code { return e.grpc }

func (e *notFoundError) HTTPStatusCode() int { return e.http }

func (e *notFoundError) APIStatusCode() string { return e.api }

type alreadyExistsError logicalError

var _ errorI = (*alreadyExistsError)(nil)

// NewAlreadyExists creates a new alreadyExistsError with corresponding HTTP
// and GPRC status codes.
func NewAlreadyExists(msg, apiStCode string) error {
	return &alreadyExistsError{
		msg:  msg,
		grpc: codes.AlreadyExists,
		http: http.StatusConflict,
		api:  apiStCode,
	}
}

func (e *alreadyExistsError) Error() string { return e.msg }

func (e *alreadyExistsError) GRPCStatusCode() codes.Code { return e.grpc }

func (e *alreadyExistsError) HTTPStatusCode() int { return e.http }

func (e *alreadyExistsError) APIStatusCode() string { return e.api }

type invalidArgumentError logicalError

var _ errorI = (*invalidArgumentError)(nil)

// NewInvalidArgument creates a new invalidArgumentError with corresponding HTTP
// and gprc status codes.
func NewInvalidArgument(msg, apiStCode string) error {
	return &invalidArgumentError{
		msg:  msg,
		grpc: codes.InvalidArgument,
		http: http.StatusBadRequest,
		api:  apiStCode,
	}
}

func (e *invalidArgumentError) Error() string { return e.msg }

func (e *invalidArgumentError) GRPCStatusCode() codes.Code { return e.grpc }

func (e *invalidArgumentError) HTTPStatusCode() int { return e.http }

func (e *invalidArgumentError) APIStatusCode() string { return e.api }
