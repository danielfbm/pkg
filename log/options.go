package log

import "io"

// SimplifiedOptions adds a simple interface to get options
type SimplifiedOptions struct {
	LogLevel     Level
	Format       EncodingFormat
	DisableColor bool
	EnableCaller bool
	Development  bool
	Writer       io.Writer `json:"-"`
}
