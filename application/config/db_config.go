package config

import (
	"go-cource-api/infrustructure/persistence"
	"os"
)

func NewDatabaseConfig() *persistence.DatabaseConfig {
	return &persistence.DatabaseConfig{
		Host: os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASSWORD"),
		User: os.Getenv("DB_USER"),
		DbName: os.Getenv("DB_DATABASE"),
		Port: os.Getenv("DB_PORT"),
	}
}