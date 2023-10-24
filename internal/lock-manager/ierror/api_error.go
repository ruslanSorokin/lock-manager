package ierror

import (
	"google.golang.org/grpc/codes"
)

type GRPCConvertible interface {
	GRPCStCode() codes.Code
}

type HTTPConvertible interface {
	HTTPStCode() int
}

type EnumConvertible interface {
	EnumStCode() string
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
func New(msg string, grpc codes.Code, http int, enum string) APIErrorI {
	return &APIError{msg: msg, grpc: grpc, http: http, enum: enum}
}

func (e *APIError) Error() string { return e.msg }

func (e *APIError) GRPCStCode() codes.Code { return e.grpc }

func (e *APIError) HTTPStCode() int { return e.http }

func (e *APIError) EnumStCode() string { return e.enum }
