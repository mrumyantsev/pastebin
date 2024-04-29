package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/errlib"
	"github.com/mrumyantsev/pastebin/internal/pkg/config"
	"github.com/mrumyantsev/pastebin/internal/pkg/endpoint"
)

type Server struct {
	config *config.Config
	server *http.Server
}

func New(cfg *config.Config, ept *endpoint.Endpoint, mv ...gin.HandlerFunc) *Server {
	if !cfg.IsEnableHttpServerDebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	ept.InitRoutes(router)

	router.Use(mv...)

	server := &http.Server{
		Addr:    cfg.HttpServerListenIp + ":" + cfg.HttpServerListenPort,
		Handler: router,
	}

	return &Server{
		config: cfg,
		server: server,
	}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil {
		return errlib.Wrap("could not start server", err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return errlib.Wrap("could not shutdown server", err)
	}

	return nil
}
