package ierror

import (
	"google.golang.org/grpc/codes"
)

type GRPCConvertible interface {
	ToGRPC() codes.Code
}

type HTTPConvertible interface {
	ToHTTP() int
}

type EnumConvertible interface {
	ToEnum() string
}

type APIErrorI interface {
	error

	GRPCConvertible
	HTTPConvertible
	EnumConvertible
}

type APIError struct {
	msg  string
	grpc codes.Code
	http int
	enum string
}

var _ APIErrorI = (*APIError)(nil)

// New creates new a APIError that meets the APIErrorI interface.
func New(msg string, grpc codes.Code, http int, enum string) *APIError {
	return &APIError{msg: msg, grpc: grpc, http: http, enum: enum}
}

func (e *APIError) Error() string { return e.msg }

func (e *APIError) ToGRPC() codes.Code { return e.grpc }

func (e *APIError) ToHTTP() int { return e.http }

func (e *APIError) ToEnum() string { return e.enum }
