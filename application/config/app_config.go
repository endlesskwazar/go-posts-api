package config

import "os"

type AppConfig struct {
	Address string
	Env     string
}

func NewAppConfig() *AppConfig {
	address :=  ":" + os.Getenv("APP_PORT")
	return &AppConfig{
		Address: address,
		Env:     os.Getenv("APP_ENV"),
	}
}