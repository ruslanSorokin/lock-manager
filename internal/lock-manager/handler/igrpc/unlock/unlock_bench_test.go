package unlock

import (
	"errors"
	"testing"

	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/provider"
	"github.com/ruslanSorokin/lock-manager/internal/lock-manager/service"
)

func BenchmarkErrToCode(b *testing.B) {
	b.Run("error=nil", benchmarkNil)
	b.Run("error=ErrInvalidResourceID", benchmarkErrInvalidResourceID)
	b.Run("error=ErrInvalidToken", benchmarkErrInvalidToken)
	b.Run("error=ErrWrongToken", benchmarkErrWrongToken)
	b.Run("error=ErrLockNotFound", benchmarkErrLockNotFound)
	b.Run("error=unexpectedError", benchmarkUnexpected)
}

func benchmarkNil(b *testing.B) {
	var e error // nil
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		errToCode(e)
	}
}

func benchmarkErrInvalidResourceID(b *testing.B) {
	e := service.ErrInvalidResourceID
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		errToCode(e)
	}
}

func benchmarkErrLockNotFound(b *testing.B) {
	e := provider.ErrLockNotFound
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		errToCode(e)
	}
}

func benchmarkErrWrongToken(b *testing.B) {
	e := provider.ErrWrongToken
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		errToCode(e)
	}
}

func benchmarkErrInvalidToken(b *testing.B) {
	e := service.ErrInvalidToken
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		errToCode(e)
	}
}

func benchmarkUnexpected(b *testing.B) {
	e := errors.New("unexpectedError")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		errToCode(e)
	}
}
