package app

import (
	"github.com/brxyxn/go-phonebook-api/config"
	"github.com/brxyxn/go-phonebook-api/pkg/logger"
)

// Configure
// This configures the settings for the app, such as the bind address.
func (a *Config) Configure() {
	vars, err := config.EnvValues() // Configuring the app variables
	if err != nil {
		logger.Error("Environment variables weren't loaded correctly!", err)
		return
	}

	a.BindAddr = ":" + vars.Server.Port
}
