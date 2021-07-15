package config

import "os"

type AppConfig struct {
	Port string
	Env string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Port: os.Getenv("APP_PORT"),
		Env: os.Getenv("APP_ENV"),
	}
}