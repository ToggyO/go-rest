package main

import (
	"fmt"
	"go-rest/internal/config"
	"go-rest/internal/infrastructure/application"
	"go-rest/internal/utils"
	"os"
)

func main() {
	envPath := utils.GetEnvFilePathFromRoot("GO_ENV", config.GO_ENV_DEVELOPMENT)
	fmt.Println(os.Getenv("GO_ENV"))

	builder := application.NewApplication(config.BuildConfigurationFromEnv(envPath))
	app, err := builder.Build()

	if err != nil {
		fmt.Println(fmt.Sprintf("Application build has failed: %s", err))
		os.Exit(1)
	}

	app.Run()
}
