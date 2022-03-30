package server

import (
	"context"
	"fmt"
	"go-rest/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler, cfg *config.Configuration) Server {
	return Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	s.check()
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	s.check()
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) check() {
	if s.httpServer == nil {
		panic("There no http server instance is provided!")
	}
}
