package config

import (
	"go-cource-api/infrustructure/persistence"
	"golang.org/x/oauth2"
)

type Config struct {
	GoogleOauthConfig   *oauth2.Config
	FaceBookOauthConfig *oauth2.Config
	DatabaseConfig      *persistence.DatabaseConfig
	AppConfig           *AppConfig
}

func NewConfig() *Config {
	appConfig := NewAppConfig()
	databaseConfig := NewDatabaseConfig()
	googleOauthConfig := NewGoogleOauthConfig()
	facebookOauthConfig := NewFaceBookOauthConfig()

	return &Config{
		GoogleOauthConfig:   googleOauthConfig,
		DatabaseConfig:      databaseConfig,
		FaceBookOauthConfig: facebookOauthConfig,
		AppConfig:           appConfig,
	}
}
