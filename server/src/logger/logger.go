package logger

import (
	"fmt"
	"os"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLogger *zap.Logger
	stdLogger *zap.SugaredLogger
)

func init() {
	var logLevel zapcore.Level
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		logLevel = zapcore.DebugLevel
		break

	case "info":
		logLevel = zapcore.InfoLevel
		break

	case "warn":
		logLevel = zapcore.WarnLevel
		break

	case "error":
		logLevel = zapcore.ErrorLevel
		break

	case "fatal":
		logLevel = zapcore.FatalLevel
		break

	default:
		logLevel = zapcore.InfoLevel
	}

	encConf := zap.NewProductionEncoderConfig()
	encConf.EncodeTime = zapcore.ISO8601TimeEncoder

	zapConf := zap.NewProductionConfig()
	zapConf.Level = zap.NewAtomicLevelAt(logLevel)
	zapConf.Encoding = "console" // "console" or "json"
	zapConf.OutputPaths = []string{"stdout"}
	zapConf.ErrorOutputPaths = []string{"stderr"}
	zapConf.EncoderConfig = encConf

	zapLogger, _ = zapConf.Build(zap.AddCallerSkip(1))
	stdLogger = zapLogger.Sugar()
}

// Debug ...
func Debug(args ...interface{}) {
	stdLogger.Debug(args...)
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	stdLogger.Debugf(format, args...)
}

// Info ...
func Info(args ...interface{}) {
	stdLogger.Info(args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	stdLogger.Infof(format, args...)
}

// Warn ...
func Warn(args ...interface{}) {
	stdLogger.Warn(args...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	stdLogger.Warnf(format, args...)
}

// Error ...
func Error(args ...interface{}) {
	stdLogger.Error(args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	stdLogger.Errorf(format, args...)
}

// Fatal ...
func Fatal(args ...interface{}) {
	stdLogger.Fatal(args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	stdLogger.Fatalf(format, args...)
}

// WarnfWithSentryCapture ...
func WarnfWithSentryCapture(format string, args ...interface{}) {
	sentry.CaptureMessage(fmt.Sprintf(format, args...))
	stdLogger.Warnf(format, args...)
}

// ErrorfWithSentryCapture ...
func ErrorfWithSentryCapture(format string, args ...interface{}) {
	sentry.CaptureMessage(fmt.Sprintf(format, args...))
	stdLogger.Errorf(format, args...)
}
