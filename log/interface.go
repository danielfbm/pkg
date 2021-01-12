package log

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testing"
)

// Logger inherit logr.Logger interface
type Logger = logr.Logger

// NullLogger null logger
type NullLogger = testing.NullLogger

// Injector allows injection of logger
type Injector interface {
	WithLogger(log Logger)
}

// MustLogger will return a NullLogger if given logger is nil
func MustLogger(logger Logger) Logger {
	if logger == nil {
		logger = NullLogger{}
	}
	return logger
}
