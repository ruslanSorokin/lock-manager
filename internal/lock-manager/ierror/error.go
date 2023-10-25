// Package provides functionality to create API-scope errors for CRUD-like
// applications.
//
// term 'static' means compile-time or immutable error such as:
//
//	var ErrAlreadyExists = ierror.NewAlreadyExists("user with this login already exists", "USER_ALREADY_EXISTS")
//
// whearas 'dynamic' means that you want to create a copy of a 'static' error
// and populate it with some information:
//
//	if (...){
//		 return nil, ierror.InstantiateAlreadyExists(ErrAlreadyExists, id)
//	}
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
