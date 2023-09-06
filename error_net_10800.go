package zapError

import (
	"go.uber.org/zap"
)

const (
	errorNetListen = 10800 + iota
)

func ErrorNetListen(network, address string, err error) []zap.Field {
	return []zap.Field{
		zap.String("network", network),
		zap.String("address", address),
		zap.Any("error", err),
		zap.Uint16("error-code", errorNetListen),
	}

}
