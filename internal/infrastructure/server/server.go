package server

import (
	"context"
	"fmt"
	"go-rest/internal/config"
	"net/http"
)

type server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler, cfg *config.Configuration) server {
	return server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Handler: handler,
		},
	}
}

func (s *server) Start() error {
	s.check()
	return s.httpServer.ListenAndServe()
}

func (s *server) Stop(ctx context.Context) error {
	s.check()
	return s.httpServer.Shutdown(ctx)
}

func (s *server) check() {
	if s.httpServer == nil {
		panic("There no http server instance is provided!")
	}
}
