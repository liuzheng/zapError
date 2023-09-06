package zapError

import "go.uber.org/zap"

const (
	ErrorCodeStrconvAtoi = 10900 + iota
	ErrorCodeStrconvParseUint
)

func ErrorStrconvAtoi(s string, err error) []zap.Field {
	return []zap.Field{
		zap.String("s", s),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeStrconvAtoi),
	}
}
func ErrorStrconvParseUint(s string, base int, bitSize int, err error) []zap.Field {
	return []zap.Field{
		zap.String("s", s),
		zap.Int("base", base),
		zap.Int("bitSize", bitSize),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeStrconvParseUint),
	}
}
