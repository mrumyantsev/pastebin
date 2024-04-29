package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/mrumyantsev/errlib"
)

// A Config is the application configuration structure.
type Config struct {
	IsEnableDebugLogs           bool `envconfig:"ENABLE_DEBUG_LOGS" default:"false"`
	IsEnableHttpServerDebugMode bool `envconfig:"ENABLE_HTTP_SERVER_DEBUG_MODE"` // default:"false"`

	HttpServerListenIp   string `envconfig:"HTTP_SERVER_LISTEN_IP" default:"0.0.0.0"`
	HttpServerListenPort string `envconfig:"HTTP_SERVER_LISTEN_PORT" default:"8080"`
}

// New creates application configuration.
func New() *Config {
	return new(Config)
}

// Init initializes application configuration.
func (c *Config) Init() error {
	if err := envconfig.Process("", c); err != nil {
		return errlib.Wrap("could not populate struct with environment variables", err)
	}

	return nil
}
