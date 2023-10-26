// Package provides functionality to create API-scope errors for CRUD-like
// applications.
//
// In this package there are 2 terms that will come up frequently:
//
// - 'static' means statically defined errors such as:
//
//	var ErrAlreadyExists = ierror.NewAlreadyExists("user with this login already exists", "USER_ALREADY_EXISTS")
//
// - 'dynamic' error is basically a copy of some 'static' error populated with
// some information:
//
//	if (...){
//	   return nil, ierror.InstantiateAlreadyExists(ErrAlreadyExists, id)
//	}
//
// All 'dynamic' errors instantiated based on some 'static' error treated as
// equal to the error they are based on if you use [errors.Is] method.
package ierror

import (
	"google.golang.org/grpc/codes"
)

type GRPCConvertible interface {
	// ToGRPC returns GRPC status code.
	ToGRPC() codes.Code
}

type HTTPConvertible interface {
	// ToHTTP returns HTTP status code in format.
	ToHTTP() int
}

type EnumConvertible interface {
	// ToEnum returns enum.
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

// New creates new a APIError.
func New(msg string, grpc codes.Code, http int, enum string) *APIError {
	return &APIError{msg: msg, grpc: grpc, http: http, enum: enum}
}

func (e *APIError) Error() string { return e.msg }

func (e *APIError) ToGRPC() codes.Code { return e.grpc }

func (e *APIError) ToHTTP() int { return e.http }

func (e *APIError) ToEnum() string { return e.enum }
