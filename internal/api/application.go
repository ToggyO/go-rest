package api

import (
	"context"
	"errors"
	api "go-rest/internal/api/http"
	"go-rest/internal/config"
	"go-rest/internal/infrastructure/ioc"
	"go-rest/internal/infrastructure/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/dig"
)

// TODO: add logger
type Application struct {
	server        *server.Server
	configuration *config.Configuration
	container     *dig.Container
}

func NewApplication(configuration *config.Configuration) (*Application, error) {
	var httpHandler http.Handler

	c, err := ioc.BuildIoc()
	err = c.Invoke(func(h http.Handler) { httpHandler = h })

	if err != nil {
		return nil, err
	}

	app := &Application{
		configuration: configuration,
		container:     c,
	}

	srv := server.NewServer(httpHandler, app.configuration)
	app.server = &srv
	return app, nil
}

func (a *Application) Run() {
	if err := a.bindRoutes(); err != nil {
		log.Fatalf("Error occurred while binding http routes: %s\n", err.Error())
	}

	go func() {
		if err := a.server.Start(); errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error occurred while running http server: %s\n", err.Error())
		}
	}()
	// TODO: log server address to console
	log.Println("Server started!")

	a.handleShutdown()
}

func (a *Application) bindRoutes() error {
	return a.container.Invoke(func(ar api.AppRouters) {
		for _, r := range ar.Routers {
			r.Bind()
		}
	})
}

func (a *Application) handleShutdown() {
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	<-quitChan

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := a.server.Stop(ctx); err != nil {
		log.Fatalf("Failed to stop server: %v", err)
	}
}
