package server

import (
	"fmt"
	"log"

	"github.com/EdwBaeza/inhouse/internal/platform/server/handler/health"
	home "github.com/EdwBaeza/inhouse/internal/platform/server/handler/home"
	postgres "github.com/EdwBaeza/inhouse/internal/platform/services/postgres"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}
	srv.registerRoutes()

	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)

	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.GET("/homes", home.ShowHandler(postgres.NewHomeRepository()))
}

func (s *Server) Engine() *gin.Engine {
	return s.engine
}
