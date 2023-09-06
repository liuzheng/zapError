package log

import (
	"syscall"

	"go.uber.org/zap"
)

const (
	ErrorCodeSyscallSignal = 11000 + iota
)

func ErrorSyscallSignal(sig syscall.Signal, err error) []zap.Field {
	return []zap.Field{
		zap.Any("sig", sig),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeSyscallSignal),
	}
}
