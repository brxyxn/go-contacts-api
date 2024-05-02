package main

import (
	server "github.com/brxyxn/go-phonebook-api/internal/app"
	"github.com/brxyxn/go-phonebook-api/internal/constants"
	"github.com/brxyxn/go-phonebook-api/pkg/logger"
	"os"
)

func main() {
	a := server.NewServer()

	// Logger configuration
	cfg := logger.LogConfig{
		Out: os.Stdout,
		LogOptions: logger.LogOptions{
			IsDebug:  true,
			LogLevel: "debug",
			CallerOptions: logger.CallerOptions{
				WithCaller:       true,
				WithCustomCaller: true,
				BasePath:         "go-phonebook-api",
			},
			DateTimeOptions: logger.DateTimeOptions{
				WithTime:   true,
				TimeFormat: constants.LongDateTimeFormat,
			},
		},
	}
	simpleLogger := logger.NewLogger(cfg)
	a.Logger = &simpleLogger

	a.Configure()

	a.Start()
}
