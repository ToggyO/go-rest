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
	GoEnv       string
	Host        string
	Port        string
	RoutePrefix string

	DbHost         string
	DbName         string
	DbUser         string
	DbPassword     string
	DbPort         string
	DbExternalPort string

	IsDev   bool
	IsStage bool
	IsProd  bool
}

func BuildConfigurationFromEnv(envFilePath string) *Configuration {
	godotenv.Load(envFilePath)
	cfg := &Configuration{
		GoEnv:       utils.GetEnv("GO_ENV", "development"),
		Host:        utils.GetEnv("GO_REST_HOST", "localhost"),
		Port:        utils.GetEnv("GO_REST_PORT", "50311"),
		RoutePrefix: utils.GetEnv("GO_REST_ROUTE_PREFIX", "50311"),

		DbHost:         utils.GetEnv("GO_REST_DB_HOST", ""),
		DbName:         utils.GetEnv("GO_REST_DB_NAME", ""),
		DbUser:         utils.GetEnv("GO_REST_DB_USER", ""),
		DbPassword:     utils.GetEnv("GO_REST_DB_PASSWORD", ""),
		DbPort:         utils.GetEnv("GO_REST_DB_PORT", ""),
		DbExternalPort: utils.GetEnv("GO_REST_DB_EXTERNAL_PORT", ""),
	}

	return buildManually(cfg)
}

func buildManually(cfg *Configuration) *Configuration {
	cfg.IsDev = cfg.GoEnv == GO_ENV_DEVELOPMENT
	cfg.IsStage = cfg.GoEnv == GO_ENV_STAGING
	cfg.IsProd = cfg.GoEnv == GO_ENV_PRODUCTION
	return cfg
}
