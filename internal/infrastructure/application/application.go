package application

import (
	"context"
	"errors"
	"go-rest/internal/application/contracts"
	"go-rest/internal/config"
	"go-rest/internal/infrastructure/ioc"
	"go-rest/internal/infrastructure/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	server        *server.Server
	configuration *config.Configuration
	container     ioc.ContainerWrapper
	logger        contracts.ILogger
}

func NewApplication(configuration *config.Configuration) (*Application, error) {
	var httpHandler http.Handler
	var logger contracts.ILogger

	// TODO: вынести создание контейнера
	c, err := ioc.NewIoc(configuration)
	err = c.GetService(func(h http.Handler) { httpHandler = h })
	err = c.GetService(func(l contracts.ILogger) { logger = l })

	if err != nil {
		return nil, err
	}

	app := &Application{
		configuration: configuration,
		container:     c,
		logger:        logger,
	}

	srv := server.NewServer(httpHandler, app.configuration)
	app.server = &srv
	return app, nil
}

func (a *Application) Run() {
	go func() {
		if err := a.server.Start(); errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error occurred while running http server: %s\n", err.Error())
		}
	}()
	// TODO: log server address to console
	log.Println("Server started!")

	a.handleShutdown()
}

func (a *Application) SetCustomIoC(ioc ioc.ContainerWrapper) {
	// TODO: implement me
	panic("implement me!")
}

func (a *Application) handleShutdown() {
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	<-quitChan

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()
	defer a.logger.Flush()

	if err := a.server.Stop(ctx); err != nil {
		log.Fatalf("Failed to stop server: %v", err)
	}
}
