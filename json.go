package logger

import (
	"os"

	"github.com/go-kit/kit/log"
)

type jsonLogger struct {
	log.Logger
}

// NewJsonLogger returns a Logger with plain cli format
func NewJsonLogger(options ...LogOption) Logger {
	opts := Option{
		w: os.Stdout,
	}

	for _, optFunc := range options {
		optFunc(&opts)
	}

	return log.NewJSONLogger(opts.w)
}
