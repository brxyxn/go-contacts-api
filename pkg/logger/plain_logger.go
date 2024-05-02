package logger

import "github.com/rs/zerolog"

var l zerolog.Logger

func Debug(message ...interface{}) {
	str := parseMessage(message...)
	l.Debug().Msg(str)
}

func Info(message ...interface{}) {
	str := parseMessage(message...)
	l.Info().Msg(str)
}

func Warn(message ...interface{}) {
	str := parseMessage(message...)
	l.Warn().Msg(str)
}

func Error(message ...interface{}) {
	str := parseMessage(message...)
	l.Error().Msg(str)
}

func Fatal(message ...interface{}) {
	str := parseMessage(message...)
	l.Fatal().Msg(str)
}

func Panic(message ...interface{}) {
	str := parseMessage(message...)
	l.Panic().Msg(str)
}
