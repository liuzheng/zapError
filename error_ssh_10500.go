package zapError

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

const (
	// crypto
	ErrorCodeCryptoSshParsePrivateKey = 10500 + iota
	ErrorCodeCryptoSshDial
	ErrorCodeCryptoSshRequestPty
	ErrorCodeCryptoSshShell
	ErrorCodeStdinPipe
	ErrorCodeStdoutPipe
)

func ErrorCryptoSshParsePrivateKey(pemBytes []byte, err error) []zap.Field {
	return []zap.Field{
		zap.Binary("pemBytes", pemBytes),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeCryptoSshParsePrivateKey),
	}
}
func ErrorCryptoSshDial(network, addr string, config *ssh.ClientConfig, err error) []zap.Field {
	return []zap.Field{
		zap.String("network", network),
		zap.String("addr", addr),
		zap.Any("config", config),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeCryptoSshDial),
	}
}
func ErrorCryptoSshRequestPty(term string, h, w int, termmodes ssh.TerminalModes, err error) []zap.Field {
	return []zap.Field{
		zap.String("term", term),
		zap.Int("h", h),
		zap.Int("w", w),
		zap.Any("termmodes", termmodes),
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeCryptoSshRequestPty),
	}
}
func ErrorCryptoSshShell(err error) []zap.Field {
	return []zap.Field{
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeCryptoSshShell),
	}
}
func ErrorStdinPipe(err error) []zap.Field {
	return []zap.Field{
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeStdinPipe),
	}
}
func ErrorStdoutPipe(err error) []zap.Field {
	return []zap.Field{
		zap.Any("error", err),
		zap.Uint16("error-code", ErrorCodeStdoutPipe),
	}
}
