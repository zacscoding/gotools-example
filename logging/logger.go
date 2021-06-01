package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

type contextKey = string

const loggerKey = contextKey("logger")

var (
	defaultLogger     *zap.SugaredLogger
	defaultLoggerOnce sync.Once
)

// NewLogger creates a new logger with the config.Context i.e config package should be initialized
func NewLogger() *zap.SugaredLogger {
	var (
		encoding = "console"
		level    = zapcore.DebugLevel
	)
	cfg := zap.Config{
		Encoding:         encoding,
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		Level:            zap.NewAtomicLevelAt(level),
		Development:      false,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := cfg.Build()
	if err != nil {
		logger = zap.NewNop()
	}
	return logger.Sugar()
}

// DefaultLogger returns the default logger for the package.
func DefaultLogger() *zap.SugaredLogger {
	defaultLoggerOnce.Do(func() {
		defaultLogger = NewLogger()
	})
	return defaultLogger
}
