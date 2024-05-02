package config

import (
	"github.com/joeshaw/envdecode"
	"time"
)

type Server struct {
	Port         string        `env:"SERVER_PORT,default=5000"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,default=5s"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,default=10s"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,default=120s"`
	IsDebug      bool          `env:"SERVER_ISDEBUG,default=false"`
	LoggerLevel  string        `env:"SERVER_LOGGER_LEVEL,default=trace"`
}

type Config struct {
	Server Server
	// here we can add more configurations
}

// EnvValues validates the current ENV_PROFILE and returns the
// values as a Config type, if ENV_PROFILE is not set the default
// environment is "dev".
func EnvValues() (*Config, error) {
	var err error

	var c Config
	err = envdecode.StrictDecode(&c)
	if err != nil {
		return nil, err
	}

	return &c, err
}
