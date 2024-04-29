package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrumyantsev/errlib"
	"github.com/mrumyantsev/pastebin/internal/pkg/config"
	"github.com/mrumyantsev/pastebin/internal/pkg/endpoint"
	"github.com/mrumyantsev/pastebin/internal/pkg/server"
	"github.com/rs/zerolog/log"
)

type App struct {
	config   *config.Config
	endpoint *endpoint.Endpoint
	server   *server.Server
}

func New() (*App, error) {
	cfg := config.New()

	err := cfg.Init()
	if err != nil {
		return nil, errlib.Wrap("could not initialize config", err)
	}

	ept := endpoint.New(cfg)

	srv := server.New(cfg, ept)

	return &App{
		config:   cfg,
		endpoint: ept,
		server:   srv,
	}, nil
}

func (a *App) Run() error {
	log.Info().Msg("service started")

	var err error // temp solution w/o db
	// err := a.database.Connect()
	// if err != nil {
	// 	return errlib.Wrap("could not connect to database", err)
	// }

	// log.Debug().Msg("database connection opened")

	if a.config.IsEnableDebugLogs {
		if a.config.IsEnableHttpServerDebugMode {
			log.Debug().Msg("starting Gin in debug mode")
		} else {
			log.Debug().Msg("starting Gin in release mode")
		}
	}

	addr := a.config.HttpServerListenIp + ":" + a.config.HttpServerListenPort

	log.Info().Msg("http server started on " + addr)

	goErr := make(chan error, 1)

	isShutdown := false

	go func() {
		if err := a.server.Start(); (err != nil) && !isShutdown {
			goErr <- errlib.Wrap("could not start http server", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-goErr:
		return err
	case <-quit:
		break
	}

	// Graceful shutdown

	log.Info().Msg("shutdown signal read")

	isShutdown = true

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	if err = a.server.Shutdown(ctx); err != nil {
		return errlib.Wrap("could not shutdown http server", err)
	}

	log.Debug().Msg("http server shut down")

	// if err = a.database.Disconnect(); err != nil {
	// 	return errlib.Wrap("could not disconnect from database", err)
	// }

	// log.Debug().Msg("database connection closed")

	log.Info().Msg("service gracefully shut down")

	return nil
}
