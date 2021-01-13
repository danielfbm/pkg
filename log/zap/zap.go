package zap

import (
	"github.com/danielfbm/pkg/log"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewDevZapLogger creates a dev zap logger
// ideal for dev environment
func NewDevZapLogger() log.Logger {
	// dev mode for now
	logger, _ := zap.NewDevelopment()
	return FromZapLogger(logger)
}

// FromZapLogger creates a logger from a *zap.Logger
func FromZapLogger(logger *zap.Logger) log.Logger {
	return zapr.NewLogger(logger)
}

// NewZapLogger returns a new logger for its options
func NewZapLogger(opts *log.SimplifiedOptions, options ...zap.Option) log.Logger {
	var config zap.Config
	if opts.Development {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	config.Level = zap.NewAtomicLevelAt(LogLevel(opts.LogLevel))

	config.Encoding = opts.Format.String()
	if !opts.DisableColor {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}
	options = append(options, zap.WithCaller(opts.EnableCaller))
	logger, _ := config.Build(options...)
	return FromZapLogger(logger)
}

// LogLevel converts string text into log level
func LogLevel(lvl log.Level) zapcore.Level {
	switch lvl {
	case log.Trace, log.Debug:
		return zapcore.DebugLevel
	case log.Info:
		return zap.InfoLevel
	case log.Warning:
		return zapcore.WarnLevel
	case log.Error:
		fallthrough
	default:
		return zapcore.ErrorLevel
	}
}
