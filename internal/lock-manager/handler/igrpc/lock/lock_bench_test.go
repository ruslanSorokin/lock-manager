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

func benchmarkErrLockAlreadyExists(b *testing.B) {
	e := provider.ErrLockAlreadyExists
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
