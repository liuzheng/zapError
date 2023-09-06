package log

import "go.uber.org/zap"

const (
	// yaml
	ErrorCodeYamlMarshal = 10200 + iota
	ErrorCodeYamlUnmarshal
)

// yaml
func ErrorYamlMarshal(in interface{}, err error) []zap.Field {
	return []zap.Field{
		zap.Any("in", in),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeYamlMarshal),
	}
}
func ErrorYamlUnmarshal(in []byte, out interface{}, err error) []zap.Field {
	return []zap.Field{
		zap.Binary("in", in),
		zap.Any("out", out),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeYamlUnmarshal),
	}
}
