package log

import (
	"os"

	"go.uber.org/zap"
)

const (
	ErrorCodeOSStat = 10300 + iota
	ErrorCodeOsOpen
	ErrorCodeOsOpenFile
	ErrorCodeOSMkdirAll
	ErrorCodeOsFindProcess
	ErrorCodeOsRemove
)

func ErrorOSStat(name string, err error) []zap.Field {
	return []zap.Field{
		zap.String("name", name),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeOSStat),
	}
}

func ErrorOsOpen(name string, err error) []zap.Field {
	return []zap.Field{
		zap.String("name", name),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeOsOpen),
	}
}
func ErrorOsOpenFile(name string, flag int, perm os.FileMode, err error) []zap.Field {
	return []zap.Field{
		zap.String("name", name),
		zap.Int("flag", flag),
		zap.Uint32("perm", uint32(perm)),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeOsOpenFile),
	}
}
func ErrorOSMkdirAll(name string, perm os.FileMode, err error) []zap.Field {
	return []zap.Field{
		zap.String("name", name),
		zap.Uint32("perm", uint32(perm)),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeOSMkdirAll),
	}
}
func ErrorOsFindProcess(pid int, err error) []zap.Field {
	return []zap.Field{
		zap.Int("pid", pid),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeOsFindProcess),
	}
}
func ErrorOsRemove(filename string, err error) []zap.Field {
	return []zap.Field{
		zap.String("filename", filename),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeOsRemove),
	}
}
