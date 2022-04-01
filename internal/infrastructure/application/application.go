package application

import (
	"context"
	"fmt"
	"go-rest/internal/application/contracts"
	"go-rest/internal/config"
	"go-rest/internal/infrastructure/ioc"
	"go-rest/internal/infrastructure/server"
	"go-rest/internal/shared/interfaces"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Application struct {
	host          interfaces.IHost
	configuration *config.Configuration
	container     interfaces.IServiceProvider
	logger        contracts.ILogger
}

func NewApplication(configuration *config.Configuration) IApplicationBuilder {
	return &Application{configuration: configuration}
}

func (a *Application) SetServiceProvider(serviceProvider interfaces.IServiceProvider) IApplicationBuilder {
	a.container = serviceProvider
	return a
}

func (a *Application) SetWebHost(host interfaces.IHost) IApplicationBuilder {
	a.host = host
	return a
}

func (a *Application) Build() (IApplication, error) {
	var container interfaces.IServiceProvider
	var httpHandler http.Handler
	var logger contracts.ILogger
	var err error

	if a.container == nil {
		container, err = ioc.NewIoc(a.configuration)
		a.container = container
	}

	if a.host == nil {
		err = a.container.GetService(func(h http.Handler) { httpHandler = h })
		srv := server.NewServer(httpHandler, a.configuration)
		a.host = &srv
	}

	err = a.container.GetService(func(l contracts.ILogger) { logger = l })

	if err != nil {
		return nil, err
	}

	a.logger = logger
	return a, err
}

func (a *Application) Run() {
	go func() {
		if err := a.host.Start(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal(fmt.Sprintf("Error occurred while running http server: %s\n", err.Error()))
		}
	}()
	// TODO: log server address to console
	a.logger.Info("Server started!")

	a.handleShutdown()
}

// TODO: user ErrorGroup package
func (a *Application) handleShutdown() {
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	<-quitChan
	a.logger.Info("Shutdown server ...")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()
	defer a.logger.Flush()

	if err := a.host.Stop(ctx); err != nil {
		a.logger.Fatal(fmt.Sprintf("Failed to stop server: %v", err))
	}

	select {
	case <-ctx.Done():
		a.logger.Info("Timeout of 5 seconds.")
	}
	a.logger.Info("Server is stopped!")
}
