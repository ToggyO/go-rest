package config

import (
	"github.com/joho/godotenv"
	"go-rest/internal/utils"
)

const (
	GO_ENV_DEVELOPMENT = "development"
	GO_ENV_STAGING     = "staging"
	GO_ENV_PRODUCTION  = "production"
)

type Configuration struct {
	GoEnv string
	Host  string
	Port  string

	IsDev   bool
	IsStage bool
	IsProd  bool
}

func BuildConfigurationFromEnv(envFilePath string) *Configuration {
	godotenv.Load(envFilePath)
	cfg := &Configuration{
		GoEnv: utils.GetEnv("GO_ENV", "development"),
		Host:  utils.GetEnv("HOST", "localhost"),
		Port:  utils.GetEnv("PORT", "50311"),
	}

	return buildManually(cfg)
}

func buildManually(cfg *Configuration) *Configuration {
	cfg.IsDev = cfg.GoEnv == GO_ENV_DEVELOPMENT
	cfg.IsStage = cfg.GoEnv == GO_ENV_STAGING
	cfg.IsProd = cfg.GoEnv == GO_ENV_PRODUCTION
	return cfg
}
