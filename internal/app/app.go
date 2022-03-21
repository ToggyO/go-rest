package app

import (
	"context"
	"errors"
	"go-rest/internal/config"
	handler "go-rest/internal/delivery/http"
	"go-rest/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TODO: add logger

type Application struct {
	server        *server.Server
	configuration *config.Configuration
}

func NewApplication(configuration *config.Configuration) *Application {
	app := &Application{configuration: configuration}

	h := handler.NewHandler().Build()
	srv := server.NewServer(h, app.configuration)

	app.server = &srv
	return app
}

func (a *Application) Run() {
	// TODO: check
	go func() {
		if err := a.server.Start(); errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error occurred while running http server: %s\n", err.Error())
		}
	}()
	log.Println("Server started!")

	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	<-quitChan

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := a.server.Stop(ctx); err != nil {
		log.Fatalf("failed to stop server: %v", err)
	}
}
