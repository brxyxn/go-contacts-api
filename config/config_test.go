package config_test

import (
	"os"
	"testing"
	"time"

	"github.com/brxyxn/go-phonebook-api/config"
	"github.com/stretchr/testify/assert"
)

func TestEnvValuesReturnsDefaultValuesWhenNoEnvSet(t *testing.T) {
	os.Clearenv()

	cfg, err := config.EnvValues()

	assert.NoError(t, err)
	assert.Equal(t, "5000", cfg.Server.Port)
	assert.Equal(t, 5*time.Second, cfg.Server.TimeoutRead)
	assert.Equal(t, 10*time.Second, cfg.Server.TimeoutWrite)
	assert.Equal(t, 120*time.Second, cfg.Server.TimeoutIdle)
	assert.Equal(t, false, cfg.Server.IsDebug)
	assert.Equal(t, "trace", cfg.Server.LoggerLevel)
}

func TestEnvValuesReturnsEnvValuesWhenSet(t *testing.T) {
	os.Setenv("SERVER_PORT", "6000")
	os.Setenv("SERVER_TIMEOUT_READ", "6s")
	os.Setenv("SERVER_TIMEOUT_WRITE", "11s")
	os.Setenv("SERVER_TIMEOUT_IDLE", "121s")
	os.Setenv("SERVER_ISDEBUG", "true")
	os.Setenv("SERVER_LOGGER_LEVEL", "debug")

	cfg, err := config.EnvValues()

	assert.NoError(t, err)
	assert.Equal(t, "6000", cfg.Server.Port)
	assert.Equal(t, 6*time.Second, cfg.Server.TimeoutRead)
	assert.Equal(t, 11*time.Second, cfg.Server.TimeoutWrite)
	assert.Equal(t, 121*time.Second, cfg.Server.TimeoutIdle)
	assert.Equal(t, true, cfg.Server.IsDebug)
	assert.Equal(t, "debug", cfg.Server.LoggerLevel)
}

func TestEnvValuesReturnsErrorWhenInvalidDuration(t *testing.T) {
	os.Setenv("SERVER_TIMEOUT_READ", "invalid")

	_, err := config.EnvValues()

	assert.Error(t, err)
}
