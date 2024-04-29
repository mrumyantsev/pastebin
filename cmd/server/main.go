package main

import (
	"os"
	"time"

	"github.com/mrumyantsev/pastebin/internal/app/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	conWrt := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
	}

	log.Logger = log.Output(conWrt)
}

// @title           Pastebin API
// @version         1.0
// @description     A system for storing blocks of text and conveniently sharing these blocks via a link.

// @contact.name   Mikhail Rumyantsev
// @contact.email  mi.rumyantsev.2020@gmail.com

// @license.name  MIT License
// @license.url   https://opensource.org/license/mit

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app, err := server.New()
	if err != nil {
		log.Fatal().Err(err).Msg("could not initialize application")
	}

	if err = app.Run(); err != nil {
		log.Fatal().Err(err).Msg("could not run application")
	}
}
