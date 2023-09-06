package log

import "go.uber.org/zap"

const (
	errorKubeNewCompleter = 10700 + iota
)

// kube
func ErrorKubeNewCompleter(err error) []zap.Field {
	return []zap.Field{
		zap.Any("error", err),
		zap.Uint16("error-code", errorKubeNewCompleter),
	}
}
