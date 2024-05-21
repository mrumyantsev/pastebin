package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/pressly/goose/v3"

	"github.com/mrumyantsev/pastebin/pkg/lib/errlib"
)

const (
	MigrationDir = "./migrations"
)

type Config struct {
	Hostname string
	Port     string
	Username string
	Password string
	Database string
	SslMode  string
	Driver   string
}

// A Database is a wrapper to sqlx.DB object to control the connection
// to a certain database.
type Database struct {
	*sqlx.DB
	config *Config
}

func New(cfg *Config) *Database {
	return &Database{config: cfg}
}

// Connect connects to the database.
func (d *Database) Connect() error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.config.Hostname,
		d.config.Port,
		d.config.Username,
		d.config.Password,
		d.config.Database,
		d.config.SslMode,
	)

	var err error

	if d.DB, err = sqlx.Open(d.config.Driver, dsn); err != nil {
		return errlib.Wrap(err, "could not connect to database")
	}

	return nil
}

// CheckConnection checks the connection to database by calling Ping
// method.
func (d *Database) CheckConnection() error {
	if err := d.Ping(); err != nil {
		return errlib.Wrap(err, "could not check the connection to database")
	}

	return nil
}

// Disconnect disconnects from the database.
func (d *Database) Disconnect() error {
	if err := d.Close(); err != nil {
		return errlib.Wrap(err, "could not disconnect from database")
	}

	return nil
}

// Migrate performs upping of the migration schema to the connected
// database, using goose migration tool.
func (d *Database) Migrate() error {
	err := goose.SetDialect(d.config.Driver)
	if err != nil {
		return errlib.Wrap(err, "could not select database dialect")
	}

	if err = goose.Up(d.DB.DB, MigrationDir); err != nil {
		return errlib.Wrap(err, "could not up the migration schema")
	}

	return nil
}
