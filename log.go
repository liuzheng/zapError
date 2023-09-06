package zapError

import (
	"os"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel int8 = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
	// None
	NoneLevel
)

var (
	logger, _ = zap.NewProduction()
	//sugar     = logger.Sugar()
	Debug  = logger.Debug
	Info   = logger.Info
	Warn   = logger.Warn
	Error  = logger.Error
	Fatal  = logger.Fatal
	DPanic = logger.DPanic
	Panic  = logger.Panic

	Binary      = zap.Binary
	Bool        = zap.Bool
	Boolp       = zap.Boolp
	ByteString  = zap.ByteString
	Complex128  = zap.Complex128
	Complex128p = zap.Complex128p
	Complex64   = zap.Complex64
	Complex64p  = zap.Complex64p
	Float64     = zap.Float64
	Float64p    = zap.Float64p
	Float32     = zap.Float32
	Float32p    = zap.Float32p
	Int         = zap.Int
	Intp        = zap.Intp
	Int64       = zap.Int64
	Int64p      = zap.Int64p
	Int32       = zap.Int32
	Int32p      = zap.Int32p
	Int16       = zap.Int16
	Int16p      = zap.Int16p
	Int8        = zap.Int8
	Int8p       = zap.Int8p
	String      = zap.String
	Strings     = zap.Strings
	Stringp     = zap.Stringp
	Uint        = zap.Uint
	Uintp       = zap.Uintp
	Uint64      = zap.Uint64
	Uint64p     = zap.Uint64p
	Uint32      = zap.Uint32
	Uint32p     = zap.Uint32p
	Uint16      = zap.Uint16
	Uint16p     = zap.Uint16p
	Uint8       = zap.Uint8
	Uint8p      = zap.Uint8p
	Uintptr     = zap.Uintptr
	Uintptrp    = zap.Uintptrp
	Reflect     = zap.Reflect
	Namespace   = zap.Namespace
	Stringer    = zap.Stringer
	Time        = zap.Time
	Timep       = zap.Timep
	Stack       = zap.Stack
	StackSkip   = zap.StackSkip
	Duration    = zap.Duration
	Durationp   = zap.Durationp
	Object      = zap.Object
	Any         = zap.Any

	humanEncoderConfig zapcore.EncoderConfig
	cores              = make(map[string]zapcore.Core)
)

type Config struct {
	Level      string `yaml:"level"`
	Path       string `yaml:"path"`
	MaxSize    int    `yaml:"max_size"`
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`
	Compress   bool   `yaml:"compress"`
}

func init() {
	humanEncoderConfig = zap.NewProductionEncoderConfig()
	humanEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	humanEncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
}
func Initial(cfg Config) {
	var level int8
	switch strings.ToLower(cfg.Level) {
	case "debug":
		level = InfoLevel // do not show debug
	case "warn", "warning":
		level = WarnLevel
	case "error":
		level = ErrorLevel
	case "dpanic":
		level = DPanicLevel
	case "panic":
		level = PanicLevel
	case "fatal":
		level = FatalLevel
	case "none":
		level = NoneLevel
	default:
		level = InfoLevel
	}
	AddFile(level, cfg)
	if level == DebugLevel {
		AddStdOut(DebugLevel)
	} else {
		AddStdOut(level)
	}
	//sugarLogger = logger.Sugar()
	createPath(cfg.Path) // todo: handle the err
	Reload(level == DebugLevel)
}
func createPath(path string) (err error) {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}
func AddStdOut(level int8) {
	if level == NoneLevel {
		delete(cores, "stdout")
	} else {
		cores["stdout"] = zapcore.NewCore(zapcore.NewConsoleEncoder(humanEncoderConfig), os.Stdout, zapcore.Level(level))
	}
}
func AddFile(level int8, cfg Config) {
	if level == NoneLevel {
		delete(cores, "file")
	} else {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   cfg.Path,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}
		cores["file"] = zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(lumberJackLogger), zapcore.Level(level))
	}

}
func Reload(AddCaller bool) {
	slice := make([]zapcore.Core, 0, len(cores))
	for _, c := range cores {
		slice = append(slice, c)
	}
	core := zapcore.NewTee(slice...)
	if AddCaller {
		logger = zap.New(core, zap.AddCaller())
	} else {
		logger = zap.New(core)
	}
	//sugarLogger = logger.Sugar()

	Debug = logger.Debug
	Info = logger.Info
	Warn = logger.Warn
	Error = logger.Error
	Fatal = logger.Fatal
	DPanic = logger.DPanic
	Panic = logger.Panic
}

func Sync() {
	logger.Sync()
}
