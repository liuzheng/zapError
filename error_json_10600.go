package log

import "go.uber.org/zap"

const (
	ErrorCodeJsonMarshal = 10600 + iota
	ErrorCodeJsonUnmarshal
)

func ErrorJsonMarshal(in interface{}, err error) []zap.Field {
	return []zap.Field{
		zap.Any("in", in),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeJsonMarshal),
	}
}
func ErrorJsonUnmarshal(data []byte, v interface{}, err error) []zap.Field {
	return []zap.Field{
		zap.Binary("data", data),
		zap.Any("v", v),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeJsonUnmarshal),
	}
}
