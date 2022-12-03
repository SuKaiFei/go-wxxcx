package log

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	_dir string
)

func init() {
	_dir = os.Getenv("LOG_DIR")
	flag.CommandLine.StringVar(&_dir, "log.dir", _dir, "log file `path, or use LOG_DIR env variable.")
}
func NewLog() log.Logger {
	logLevel := "info"

	if len(_dir) == 0 {
		_dir = "./logs/app.log"
	}

	hook := lumberjack.Logger{
		Filename:   _dir,
		MaxSize:    200,
		MaxAge:     6,
		MaxBackups: 20,
		Compress:   false,
	}

	level := zap.InfoLevel
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
	logger := NewZapLogger(zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.AddSync(&hook),
			level,
		),
		zap.AddCaller(),
		zap.AddCallerSkip(3),
		zap.AddStacktrace(zap.ErrorLevel),
	))
	log.DefaultLogger = logger

	return logger
}
