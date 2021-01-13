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

// Level as an integer, used mostly
type Level int

// LevelStr levels as strings
type LevelStr string

const (
	T        int      = 3
	Trace    Level    = 3
	TraceStr LevelStr = "trace"

	D        int      = 2
	Debug    Level    = 2
	DebugStr LevelStr = "debug"

	I       int      = 1
	Info    Level    = 1
	InfoStr LevelStr = "info"

	W          int      = 0
	Warning    Level    = Level(W)
	WarningStr LevelStr = "warning"

	E        int      = -1
	Error    Level    = -1
	ErrorStr LevelStr = "error"
)

// Level level from string
func (s LevelStr) Level() Level {
	switch s {
	case TraceStr:
		return Trace
	case DebugStr:
		return Debug
	case InfoStr:
		return Info
	case WarningStr:
		return Warning
	case ErrorStr:
		fallthrough
	default:
		return Error
	}
}

// LevelStr as string
func (s Level) LevelStr() LevelStr {
	switch s {
	case Trace:
		return TraceStr
	case Debug:
		return DebugStr
	case Info:
		return InfoStr
	case Warning:
		return WarningStr
	case Error:
		fallthrough
	default:
		return ErrorStr
	}
}

// Int return value as int
func (s Level) Int() int {
	return int(s)
}

// String return itself as string
func (s LevelStr) String() string {
	return string(s)
}

// EncodingFormat format for encoding
type EncodingFormat string

const (
	ConsoleFormat EncodingFormat = "console"
	JSONFormat    EncodingFormat = "json"
)

// String retruns as string
func (f EncodingFormat) String() string {
	return string(f)
}
