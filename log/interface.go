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
	Trace    Level    = 3
	T        Level    = Trace
	TraceStr LevelStr = "trace"

	Debug    Level    = 2
	D        Level    = Debug
	DebugStr LevelStr = "debug"

	Info    Level    = 1
	I       Level    = Info
	InfoStr LevelStr = "info"

	Warning    Level    = 0
	W          Level    = Warning
	WarningStr LevelStr = "warning"

	Error    Level    = -1
	E        Level    = Error
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
