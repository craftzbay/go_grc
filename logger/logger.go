package logger

import (
	"fmt"
	"net/url"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	// DPanic, Panic and Fatal level can not be set by user
	DebugLevelStr   string = "debug"
	InfoLevelStr    string = "info"
	WarningLevelStr string = "warning"
	ErrorLevelStr   string = "error"
)

var (
	globalLogger *zap.Logger
)

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}

func InitLogger(logLevel string, logFile string, isProd bool) error {

	var level zapcore.Level
	switch logLevel {
	case DebugLevelStr:
		level = zap.DebugLevel
	case InfoLevelStr:
		level = zap.InfoLevel
	case WarningLevelStr:
		level = zap.WarnLevel
	case ErrorLevelStr:
		level = zap.ErrorLevel
	default:
		return fmt.Errorf("unknown log level %s", logLevel)
	}

	writer := lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    1024, //MB
		MaxBackups: 30,
		MaxAge:     90, //days
		Compress:   true,
	}
	zap.RegisterSink("lumberjack", func(*url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &writer,
		}, nil
	})
	//config := zap.NewProductionEncoderConfig()
	config := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(&writer), level),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level),
	)

	var _globalLogger *zap.Logger
	if isProd {
		_globalLogger = zap.New(core, zap.AddCaller())
	} else {
		_globalLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(level))
	}
	zap.ReplaceGlobals(_globalLogger)
	globalLogger = _globalLogger
	return nil
}

func Info(message string, fields ...zap.Field) {
	globalLogger.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	globalLogger.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	globalLogger.Error(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	globalLogger.Warn(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	globalLogger.Fatal(message, fields...)
}

func NewSugar(name string) *zap.SugaredLogger {
	return globalLogger.Named(name).Sugar()
}
