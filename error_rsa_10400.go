package zapError

import (
	"io"

	"go.uber.org/zap"
)

const (
	errorRsaGenerateKey = 10400 + iota
)

func ErrorRsaGenerateKey(random io.Reader, bits int, err error) []zap.Field {
	return []zap.Field{
		zap.Any("random", random),
		zap.Int("bits", bits),
		zap.Any("error", err),
		zap.Uint16("error-code", errorRsaGenerateKey),
	}
}
