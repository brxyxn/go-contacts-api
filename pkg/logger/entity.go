package logger

import (
	"github.com/rs/zerolog"
	"io"
)

var (
	notFound     = -1 // Not found index
	DefaultIndex = 1  // Use when you want to display the caller from the next level after the basePath
)

// SimpleLogger is the simplest way to log a message at the
// given level this logger uses zerolog to log the message.
type SimpleLogger interface {
	Debug() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
	Error() *zerolog.Event
	Fatal() *zerolog.Event
	Panic() *zerolog.Event
	Print(message ...interface{})
	Printf(format string, message ...interface{})
}

// ZerologAdapter is the standard interface for logging using the zerolog basic methods/options
type ZerologAdapter interface {
	Trace() *zerolog.Event
	Debug() *zerolog.Event
	Info() *zerolog.Event
	Warn() *zerolog.Event
	Error() *zerolog.Event
	Fatal() *zerolog.Event
	Panic() *zerolog.Event
	Log() *zerolog.Event
	Err(err error) *zerolog.Event
	WithLevel(level zerolog.Level) *zerolog.Event
	With() zerolog.Context
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Output(level zerolog.Level) zerolog.Logger
}

// Logger is the standard interface for Simple Logger interface
type Logger struct {
	logger *zerolog.Logger
}

// ZeroLogger is the standard interface for logging using the zerolog basic methods/options
type ZeroLogger struct {
	logger *zerolog.Logger
}

// LogConfig contains the properties to configure the logger
type LogConfig struct {
	Out io.Writer

	LogOptions
}

type LogOptions struct {
	// If IsDebug is set to true the property LogLevel will be overrided with Trace Level for debugging
	// as TraceLevel is the lowest log level accepted
	IsDebug bool
	// LogLevel accepts zerlog.Level*Value values only, example:
	// "trace", "debug", "info", "warn", "error", "fatal", "panic"
	LogLevel string

	CallerOptions
	ConsoleOptions
	DateTimeOptions
}

// CallerOptions are options to customize caller field in logging
type CallerOptions struct {
	WithCaller       bool
	WithCustomCaller bool
	BasePath         string
}

// ConsoleOptions are options to customize the output and configure pretty console output, take into considerations
// that prettyfing the output will cost processing the time and resources used for logging
type ConsoleOptions struct {
	IsConsole     bool
	WithoutColor  bool
	FormatMessage string
}

type DateTimeOptions struct {
	WithTime   bool
	TimeFormat string
}
