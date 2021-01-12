package zap

import (
	"github.com/danielfbm/pkg/log"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
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
