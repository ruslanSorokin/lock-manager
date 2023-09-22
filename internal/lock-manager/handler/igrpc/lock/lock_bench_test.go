package lock

import (
	"errors"
	"testing"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
)

func BenchmarkErrToCode(b *testing.B) {
	b.Run("error=nil", benchmarkNil)
	b.Run("error=ErrInvalidResourceID", benchmarkErrInvalidResourceID)
	b.Run("error=ErrLockAlreadyExists", benchmarkErrLockAlreadyExists)
	b.Run("error=unexpectedError", benchmarkUnexpected)
}

func benchmarkNil(b *testing.B) {
	var e error // nil
	msgMapper := newErrToCodeMapper()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgMapper(e)
	}
}

func benchmarkErrInvalidResourceID(b *testing.B) {
	e := service.ErrInvalidResourceID
	msgMapper := newErrToCodeMapper()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgMapper(e)
	}
}

func benchmarkErrLockAlreadyExists(b *testing.B) {
	e := provider.ErrLockAlreadyExists
	msgMapper := newErrToCodeMapper()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgMapper(e)
	}
}

func benchmarkUnexpected(b *testing.B) {
	e := errors.New("unexpectedError")
	msgMapper := newErrToCodeMapper()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		msgMapper(e)
	}
}
