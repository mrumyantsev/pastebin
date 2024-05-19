package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/pastebin/pkg/lib/errlib"
)

type RoutesInitilizer interface {
	InitRoutes(engine *gin.Engine)
}

type Config struct {
	ListenIpAddress string
	ListenPort      string

	// Maximum headers with its values to parse in requests. If passed
	// zero it will be 1 << 20 (1 megabyte).
	MaxHeaderBytes int

	// Timeout for reading the requests. If passed zero it will be 15
	// (15 seconds).
	ReadTimeoutSeconds int

	// Timeout for writing the responses. If passed zero it will be 15
	// (15 seconds).
	WriteTimeoutSeconds int
}

func (c *Config) init() {
	if c.MaxHeaderBytes == 0 {
		c.MaxHeaderBytes = 1 << 20 // 1 megabyte
	}
	if c.ReadTimeoutSeconds == 0 {
		c.ReadTimeoutSeconds = 15 // 15 seconds
	}
	if c.WriteTimeoutSeconds == 0 {
		c.WriteTimeoutSeconds = 15 // 15 seconds
	}
}

type Server struct {
	config *Config
	engine *gin.Engine
	server *http.Server
}

func New(cfg *Config) *Server {
	cfg.init()

	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	server := &http.Server{
		Addr:           cfg.ListenIpAddress + ":" + cfg.ListenPort,
		Handler:        engine,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		ReadTimeout:    time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
	}

	return &Server{
		config: cfg,
		engine: engine,
		server: server,
	}
}

func (s *Server) AddMW(mw ...gin.HandlerFunc) {
	s.engine.Use(mw...)
}

func (s *Server) InitRoutes(endpoints ...RoutesInitilizer) {
	for _, endpoint := range endpoints {
		endpoint.InitRoutes(s.engine)
	}
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil {
		return errlib.Wrap(err, "could not proceed serving")
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return errlib.Wrap(err, "could not shutdown")
	}

	return nil
}
