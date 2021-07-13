package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/application/config"
	"go-cource-api/application/handlers"
	"go-cource-api/application/middlewares"
	"go-cource-api/infrustructure/persistence"
	"go-cource-api/infrustructure/security"
	"go-cource-api/routes"
)

func main() {
	appConfig := config.NewConfig()

	services, err := persistence.NewRepositories(
		appConfig.DatabaseConfig.User,
		appConfig.DatabaseConfig.Password,
		appConfig.DatabaseConfig.Port,
		appConfig.DatabaseConfig.Host,
		appConfig.DatabaseConfig.DbName,
	)

	if err != nil {
		panic(err)
	}

	if err = services.Automigrate(); err != nil {
		panic(err)
	}

	postHandlers := handlers.NewPosts(services.Post)
	commentHandlers := handlers.NewComments(services.Comment)
	securityHandlers := handlers.NewSecurity(security.NewSecurity(services.User))

	responseResponder := application.NewResponseResponder()

	e := echo.New()
	e.Use(middlewares.ConfigInjectorMiddleware(appConfig))
	e.Use(middlewares.ResponderInjectorMiddleware(responseResponder))
	e.Validator = &application.CustomValidator{
		Validator: validator.New(),
	}
	e.Renderer = application.Renderer()

	routes.InitAuthRoutes(e, securityHandlers)
	routes.InitApiV1Routes(e, postHandlers, commentHandlers)

	// Start server
	serverUrl := ":" + appConfig.AppConfig.Port
	e.Logger.Fatal(e.Start(serverUrl))
}
