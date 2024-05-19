package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrumyantsev/pastebin/internal/domain/user"
	"github.com/mrumyantsev/pastebin/internal/pkg/config"
	"github.com/mrumyantsev/pastebin/internal/pkg/database"
	"github.com/mrumyantsev/pastebin/internal/pkg/httpserver"
	"github.com/mrumyantsev/pastebin/pkg/lib/errlib"
	"github.com/rs/zerolog/log"
)

const (
	httpServerShutdownTimeoutSeconds = 5
)

type UseCase struct {
	userUseCase user.UserUseCase
}

type App struct {
	config     *config.Config
	database   *database.Database
	httpServer *httpserver.Server
	useCase    *UseCase
	isShutdown bool
}

func New() (*App, error) {
	app := &App{}

	err := app.InitConfig()
	if err != nil {
		return nil, errlib.Wrap(err, "could not initialize configuration")
	}

	app.InitDb()

	app.InitHttpServer()

	app.InitUseCase()

	app.InitHttpEndpoints()

	return app, nil
}

func (a *App) Run() error {
	errCh := make(chan error)
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, syscall.SIGINT, syscall.SIGTERM)

	err := a.connectToDb()
	if err != nil {
		return err
	}

	go a.startHttpServer(errCh)

	select {
	case err = <-errCh:
		return err
	case <-exitCh:
		a.isShutdown = true
	}

	if err = a.stopHttpServer(); err != nil {
		return err
	}

	if err = a.disconnectFromDb(); err != nil {
		return err
	}

	return nil
}

func (a *App) connectToDb() error {
	if err := a.database.Connect(); err != nil {
		return errlib.Wrap(err, "could not connect to database")
	}

	log.Debug().Msg("database connected")

	return nil
}

func (a *App) disconnectFromDb() error {
	if err := a.database.Disconnect(); err != nil {
		return errlib.Wrap(err, "could not disconnect from database")
	}

	log.Debug().Msg("database disconnected")

	return nil
}

func (a *App) startHttpServer(errCh chan error) {
	addr := a.config.HttpServerListenIpAddress + ":" + a.config.HttpServerListenPort

	log.Info().Msg("http server started on " + addr)

	if err := a.httpServer.Start(); (err != nil) && !a.isShutdown {
		errCh <- err
	}
}

func (a *App) stopHttpServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(httpServerShutdownTimeoutSeconds)*time.Second)
	defer cancel()

	if err := a.httpServer.Shutdown(ctx); err != nil {
		return errlib.Wrap(err, "could not shutdown the http server")
	}

	log.Debug().Msg("http server shut down")

	return nil
}
