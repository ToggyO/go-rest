package config

import (
	"github.com/joho/godotenv"
	"go-rest/internal/utils"
)

type Configuration struct {
	Host string
	Port string
}

func BuildConfigurationFromEnv(envFilePath string) *Configuration {
	godotenv.Load(envFilePath)
	return &Configuration{
		Host: utils.GetEnv("HOST", "localhost"),
		Port: utils.GetEnv("PORT", "50311"),
	}
}
