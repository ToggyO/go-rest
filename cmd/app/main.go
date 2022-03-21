package main

import (
	"go-rest/internal/app"
	"go-rest/internal/config"
)

func main() {
	application := app.NewApplication(
		config.BuildConfigurationFromEnv("../../.env"),
	)
	application.Run()
}
