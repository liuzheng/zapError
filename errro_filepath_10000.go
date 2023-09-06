package zapError

import "go.uber.org/zap"

const (
	errorFilepathAbs = 10000 + iota
)

func ErrorFilepathAbs(path string, err error) []zap.Field {
	return []zap.Field{
		zap.String("path", path),
		zap.Any("error", err),
		zap.Uint16("error-code", errorFilepathAbs),
	}
}
