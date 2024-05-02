package app

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/brxyxn/go-phonebook-api/pkg/logger"
)

type Config struct {
	Ctx        context.Context
	Logger     *logger.SimpleLogger
	Router     *mux.Router
	BindAddr   string
	httpServer *http.Server
}

// NewServer ...
func NewServer() *Config {
	router := mux.NewRouter()
	return &Config{Router: router}
}

// Start ...
func (a *Config) Start() {
	// Creating a new server
	a.httpServer = &http.Server{
		Addr:         a.BindAddr,        // configure the bind address
		Handler:      a.Router,          // set the default handler
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		// ErrorLog:     a.Logger,            // set the logger for the server
	}

	a.setupRoutes()

	// Starting the server
	go func() {
		logger.Info("running server on port", a.BindAddr)

		err := a.httpServer.ListenAndServe()
		if err != nil {
			logger.Info("Server Status: ", err)
			os.Exit(1)
		}
	}()

	// Creating channel
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt, os.Kill)
	// signal.Notify(cs, os.Kill) // If running on Windows

	sigchan := <-cs
	logger.Debug("Signal received:", sigchan)

	ctx, fn := context.WithTimeout(context.Background(), 30*time.Second)
	defer fn()
	err := a.httpServer.Shutdown(ctx)
	if err != nil {
		return
	}
}
