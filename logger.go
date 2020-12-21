package logger

import (
	"fmt"
	"io"
)

type Option struct {
	w io.Writer
}

type LogOption func(*Option)

var errMissingValue = fmt.Errorf("(MISSING)")

type Logger interface {
	Log(keyvals ...interface{}) error
}

// WithOutput configures the logger output to w io.Writer
func WithOutput(w io.Writer) LogOption {
	return func(o *Option) {
		o.w = w
	}
}

// LoggerFunc is an adapter to allow use of ordinary functions as Loggers. If
// f is a function with the appropriate signature, LoggerFunc(f) is a Logger
// object that calls f.
type LoggerFunc func(keyvals ...interface{}) error

// Log implements Logger by calling f(keyvals...).
func (f LoggerFunc) Log(keyvals ...interface{}) error {
	return f(keyvals...)
}
