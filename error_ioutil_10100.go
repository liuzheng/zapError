package zapError

import (
	"io"
	"os"

	"go.uber.org/zap"
)

const (
	// ioutil
	errorioutilReadFile = 10100 + iota
	errorioutilWriteFile
)

// ioutil
func ErrorIoutilReadFile(filename string, err error) []zap.Field {
	return []zap.Field{
		zap.String("filename", filename),
		zap.Any("error", err),
		zap.Uint16("error-code", errorioutilReadFile),
	}
}
func ErrorIoutilReadAll(r io.Reader, err error) []zap.Field {
	return []zap.Field{
		zap.Any("r", r),
		zap.Any("error", err),
		zap.Uint16("error-code", errorioutilReadFile),
	}
}
func ErrorIoutilWriteFile(filename string, data []byte, perm os.FileMode, err error) []zap.Field {
	return []zap.Field{
		zap.String("filename", filename),
		zap.Binary("data", data),
		zap.Uint32("perm", uint32(perm)),
		zap.Any("error", err),
		zap.Uint16("error-code", errorioutilWriteFile),
	}
}
