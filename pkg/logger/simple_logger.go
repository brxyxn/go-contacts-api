package logger

import (
	"github.com/rs/zerolog"
)

func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}

func (l *Logger) Panic() *zerolog.Event {
	return l.logger.Panic()
}

func (l *Logger) Print(message ...interface{}) {
	l.logger.Print(message...)
}

func (l *Logger) Printf(format string, message ...interface{}) {
	l.logger.Printf(format, message...)
}
