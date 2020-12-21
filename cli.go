package logger

import (
	"fmt"
	"io"
	"os"

	apexlog "github.com/apex/log"
	apexcli "github.com/apex/log/handlers/cli"
)

type cliLogger struct {
	field  *apexlog.Logger
	level  apexlog.Level
	msgKey string
}

// NewCliLogger returns a Logger with plain cli format
func NewCliLogger(msgKey string, options ...LogOption) Logger {
	opts := Option{
		w: os.Stdout,
	}

	for _, optFunc := range options {
		optFunc(&opts)
	}

	fieldLogger := &apexlog.Logger{}
	fieldLogger.Handler = apexcli.New(opts.w)
	fieldLogger.Level = apexlog.DebugLevel
	l := &cliLogger{
		field:  fieldLogger,
		level:  apexlog.InfoLevel,
		msgKey: msgKey,
	}

	return l
}

func (l cliLogger) Log(keyvals ...interface{}) error {
	fields := apexlog.Fields{}
	var msg = ""
	for i := 0; i < len(keyvals); i += 2 {
		if keyvals[i] == l.msgKey {
			if i+1 < len(keyvals) {
				msg = fmt.Sprintf("%v", keyvals[i+1])
			}
			continue
		}
		if i+1 < len(keyvals) {
			fields[fmt.Sprint(keyvals[i])] = keyvals[i+1]
		} else {
			fields[fmt.Sprint(keyvals[i])] = errMissingValue
		}
	}

	switch l.level {
	case apexlog.InfoLevel:
		l.field.WithFields(fields).Info(msg)
	case apexlog.ErrorLevel:
		l.field.WithFields(fields).Error(msg)
	case apexlog.DebugLevel:
		l.field.WithFields(fields).Debug(msg)
	case apexlog.WarnLevel:
		l.field.WithFields(fields).Warn(msg)
	default:
		l.field.WithFields(fields).Trace(msg)
	}

	return nil
}

func (l *cliLogger) SetOutput(out io.Writer) {
	l.field.Handler = apexcli.New(out)
}
