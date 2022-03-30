package main

import (
	"fmt"
	"go-rest/internal/config"
	"go-rest/internal/infrastructure/application"
	"os"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()
	envPath := filepath.Join(pwd, ".env")

	app, err := application.NewApplication(
		config.BuildConfigurationFromEnv(envPath),
	)
	// TODO: handle error properly
	if err != nil {
		panic(fmt.Sprintf("Failed to start application: %s", err))
	}

	app.Run()
}
