// Package logging defines the minimal interface that loggers must support to be used by pack.
package logging

import (
	"io"

	"github.com/buildpack/lifecycle/style"
)

const (
	InfoLevel  = "info"
	DebugLevel = "debug"
	WarnLevel  = "warn"
)

// Logger defines behavior required by a logging package used by pack libraries
type Logger interface {
	Debug(msg string)
	Debugf(fmt string, v ...interface{})

	Info(msg string)
	Infof(fmt string, v ...interface{})

	Warn(msg string)
	Warnf(fmt string, v ...interface{})

	Error(msg string)
	Errorf(fmt string, v ...interface{})

	Writer() io.Writer

	WantLevel(level string)
}

// WithDebugErrorWriter is an optional interface for loggers that want to support a separate writer for errors and standard logging.
// the DebugErrorWriter should write to stderr if quiet is false.
type WithDebugErrorWriter interface {
	DebugErrorWriter() io.Writer
}

// WithDebugWriter is an optional interface what will return a writer that will write raw output if quiet is false.
type WithDebugWriter interface {
	DebugWriter() io.Writer
}

// GetDebugErrorWriter will return an ErrorWriter, typically stderr if one exists, otherwise the standard logger writer
// will be returned.
func GetDebugErrorWriter(l Logger) io.Writer {
	if er, ok := l.(WithDebugErrorWriter); ok {
		return er.DebugErrorWriter()
	}
	return l.Writer()
}

// GetDebugWriter returns a writer
// See WithDebugWriter
func GetDebugWriter(l Logger) io.Writer {
	if ew, ok := l.(WithDebugWriter); ok {
		return ew.DebugWriter()
	}
	return l.Writer()
}

// Tip logs a tip.
func Tip(l Logger, format string, v ...interface{}) {
	l.Infof(style.Tip("Tip: ")+format, v...)
}
