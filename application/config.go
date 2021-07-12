package application

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"os"
)

type DatabaseConfig struct {
	Host string
	Password string
	User string
	DbName string
	Port string
}

type Config struct {
	GoogleOauthConfig *oauth2.Config
	FaceBookOauthConfig *oauth2.Config
	DatabaseConfig *DatabaseConfig
}

func NewConfig() *Config {
	databaseConfig := &DatabaseConfig{
		Host: os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASSWORD"),
		User: os.Getenv("DB_USER"),
		DbName: os.Getenv("DB_DATABASE"),
		Port: os.Getenv("DB_PORT"),
	}

	googleOauthConfig := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_SUCCESS_URL"),
		Scopes: []string{
			"profile",
			"email",
		},
		Endpoint: google.Endpoint,
	}

	facebookOauthConfig := &oauth2.Config{
		ClientID: os.Getenv("FACEBOOK_CLIENT_ID"),
		ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
		RedirectURL: os.Getenv("FACEBOOK_SUCCESS_URL"),
		Scopes: []string{
			"public_profile",
			"email",
		},
		Endpoint: facebook.Endpoint,
	}

	return &Config{
		GoogleOauthConfig: googleOauthConfig,
		DatabaseConfig: databaseConfig,
		FaceBookOauthConfig: facebookOauthConfig,
	}
}
