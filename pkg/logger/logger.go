package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

func init() {
	l = zerolog.New(os.Stdout)
}

// NewLogger creates a new logger instance
func NewLogger(cfg LogConfig) SimpleLogger {
	return &Logger{buildLogger(cfg)}
}

// buildLogger builds a logger instance based on the given config and returns a zerolog.Logger instance
func buildLogger(cfg LogConfig) *zerolog.Logger {
	// custom message key identifier
	zerolog.MessageFieldName = "msg"

	// validating if logger will print custom caller
	if cfg.CallerOptions.WithCustomCaller && cfg.CallerOptions.BasePath != "" {
		zerolog.CallerMarshalFunc = customCallerMarshalFunc(cfg.BasePath)
	}

	// setting zerolog.LogLevel
	logLevel := getLogLevel(cfg.LogLevel)
	if cfg.IsDebug {
		logLevel = zerolog.TraceLevel
	}
	zerolog.SetGlobalLevel(logLevel)

	// Create logger instance
	logger := zerolog.Logger{}

	// if IsConsole then format Out value for console style
	if cfg.ConsoleOptions.IsConsole {
		out := zerolog.NewConsoleWriter()
		cfg.Out = console(out, cfg)
	}

	logger = zerolog.New(cfg.Out)

	// creating context for logger options
	ctx := logger.With()

	// if DateTime is enabled then adding DateTime Field to the logger
	if cfg.DateTimeOptions.WithTime {
		ctx = ctx.Timestamp()
	}

	// if Caller is enabled then adding Caller Field to the logger
	if cfg.CallerOptions.WithCaller {
		ctx = ctx.Caller()
	}

	// adding options to logger instance
	logger = ctx.Logger()

	return &logger
}

// console returns instance for zerolog.ConsoleWriter type.
func console(out zerolog.ConsoleWriter, cfg LogConfig) zerolog.ConsoleWriter {
	out.NoColor = cfg.WithoutColor
	out.TimeFormat = cfg.TimeFormat
	out.FormatMessage = format(cfg.FormatMessage)
	return out
}

// format returns the zerolog.Formatter type for formatting the message on the console output.
func format(value interface{}) zerolog.Formatter {
	return func(i interface{}) string {
		var str string
		if s, ok := value.(string); ok {
			str = s
		}
		return str + " " + fmt.Sprint(i)
	}
}

// parseMessage converts the input of type array into a string to be printed out from the logger method.
func parseMessage(message ...interface{}) string {
	str := ""
	lenght := len(message) - 1
	for i, s := range message {
		if i == lenght {
			str = str + fmt.Sprintf("%v", s)
			continue
		}
		str = str + fmt.Sprintf("%v ", s)
	}
	return str
}

// getLogLevel returns the zerolog.*Level corresponding to the
// given log level as a string it expects a zerolog.Level* input.
func getLogLevel(loglevel string) zerolog.Level {
	switch loglevel {
	case zerolog.LevelTraceValue:
		return zerolog.TraceLevel

	case zerolog.LevelDebugValue:
		return zerolog.DebugLevel

	case zerolog.LevelInfoValue:
		return zerolog.InfoLevel

	case zerolog.LevelWarnValue:
		return zerolog.WarnLevel

	case zerolog.LevelErrorValue:
		return zerolog.ErrorLevel

	case zerolog.LevelFatalValue:
		return zerolog.FatalLevel

	case zerolog.LevelPanicValue:
		return zerolog.PanicLevel

	case "disabled":
		return zerolog.Disabled

	case "nolevel":
		return zerolog.NoLevel

	default:
		return zerolog.InfoLevel
	}
}
