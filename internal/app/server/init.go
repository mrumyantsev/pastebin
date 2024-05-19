package server

import (
	v1 "github.com/mrumyantsev/pastebin/internal/domain/user/endpoint/http/v1"
	"github.com/mrumyantsev/pastebin/internal/domain/user/repository/postgres"
	"github.com/mrumyantsev/pastebin/internal/domain/user/service"
	"github.com/mrumyantsev/pastebin/internal/domain/user/usecase"
	"github.com/mrumyantsev/pastebin/internal/pkg/config"
	"github.com/mrumyantsev/pastebin/internal/pkg/database"
	"github.com/mrumyantsev/pastebin/internal/pkg/httpserver"
	"github.com/mrumyantsev/pastebin/pkg/lib/errlib"
)

func (a *App) InitConfig() error {
	a.config = config.New()

	if err := a.config.Init(); err != nil {
		return errlib.Wrap(err, "could not complete init")
	}

	return nil
}

func (a *App) InitDb() {
	dbCfg := &database.Config{
		Hostname: a.config.DbHostname,
		Port:     a.config.DbPort,
		Username: a.config.DbUsername,
		Password: a.config.DbPassword,
		Database: a.config.DbDatabase,
		SslMode:  a.config.DbSslMode,
		Driver:   a.config.DbDriver,
	}

	a.database = database.New(dbCfg)
}

func (a *App) InitHttpServer() {
	hsrvCfg := &httpserver.Config{
		ListenIpAddress: a.config.HttpServerListenIpAddress,
		ListenPort:      a.config.HttpServerListenPort,
	}

	a.httpServer = httpserver.New(hsrvCfg)
}

func (a *App) InitUseCase() {
	userRepo := postgres.NewUserPostgresRepository(a.database)

	userService := service.NewUserService(userRepo)

	a.useCase = &UseCase{
		userUseCase: usecase.NewUserUseCase(a.config, userService),
	}
}

func (a *App) InitHttpEndpoints() {
	userEndpoint := v1.NewUserV1HttpEndpoint(a.config, a.useCase.userUseCase)

	a.httpServer.InitRoutes(userEndpoint)
}
