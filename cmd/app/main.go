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

	builder := application.NewApplication(config.BuildConfigurationFromEnv(envPath))
	app, err := builder.Build()

	if err != nil {
		fmt.Println(fmt.Sprintf("Application build has failed: %s", err))
		os.Exit(1)
	}

	app.Run()
}
