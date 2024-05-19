package config

import (
	"errors"

	"github.com/kelseyhightower/envconfig"
	"github.com/mrumyantsev/pastebin/pkg/lib/errlib"
)

const (
	EnvPrefix = ""
)

// A Config is the application configuration structure.
type Config struct {
	DbHostname string `envconfig:"DB_HOSTNAME" default:"localhost"`
	DbPort     string `envconfig:"DB_PORT" default:"5432"`
	DbUsername string `envconfig:"DB_USERNAME" default:"postgres"`
	DbPassword string `envconfig:"DB_PASSWORD"`
	DbDatabase string `envconfig:"DB_DATABASE" default:"pastebin"`
	DbSslMode  string `envconfig:"DB_SSLMODE" default:"disable"`
	DbDriver   string `envconfig:"DB_DRIVER" default:"postgres"`

	HttpServerListenIpAddress string `envconfig:"HTTP_SERVER_LISTEN_IP_ADDRESS" default:"0.0.0.0"`
	HttpServerListenPort      string `envconfig:"HTTP_SERVER_LISTEN_PORT" default:"8080"`

	PasswordHashSalt string `envconfig:"PASSWORD_HASH_SALT"`

	ItemsOnPage    int `envconfig:"ITEMS_ON_PAGE" default:"2"`
	MaxItemsOnPage int `envconfig:"MAX_ITEMS_ON_PAGE" default:"4"`
}

// New creates an application configuration.
func New() *Config {
	return &Config{}
}

// Init initializes application configuration.
func (c *Config) Init() error {
	if err := envconfig.Process(EnvPrefix, c); err != nil {
		return errlib.Wrap(err, "could not populate config structure")
	}

	if c.DbPassword == "" {
		return errors.New("no database password specified")
	}

	if c.PasswordHashSalt == "" {
		return errors.New("no password hash salt specified")
	}

	return nil
}
