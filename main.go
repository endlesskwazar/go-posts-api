package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"go-cource-api/application"
	"go-cource-api/application/config"
	"go-cource-api/application/handlers"
	"go-cource-api/application/middlewares"
	_ "go-cource-api/docs"
	"go-cource-api/infrustructure/persistence"
	"go-cource-api/infrustructure/services"
	"go-cource-api/routes"
)

// @title Posts API documentation
// @version 1.0
// @description Swagger API for Golang Post Project.
// @host http://localhost:8000

// @contact.name Alexandr
// @contact.email endlesskwazar@gmail.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	appConfig := config.NewConfig()

	repositories, err := persistence.NewRepositories(appConfig.DatabaseConfig)

	if err != nil {
		panic(err)
	}

	if err = repositories.Automigrate(); err != nil {
		panic(err)
	}

	postHandlers := handlers.NewPostHandlers(repositories.Post)
	commentHandlers := handlers.NewCommentHandlers(repositories.Comment)
	securityHandlers := handlers.NewSecurityHandlers(services.NewSecurityService(repositories.User))

	responder := application.NewResponseResponder()

	e := echo.New()
	e.Use(middlewares.ConfigInjectorMiddleware(appConfig))
	e.Use(middlewares.ResponderInjectorMiddleware(responder))
	e.Validator = &application.CustomValidator{
		Validator: validator.New(),
	}
	e.HTTPErrorHandler = application.ErrorHandler
	e.Renderer = application.Renderer()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.InitAuthRoutes(e, securityHandlers)
	routes.InitApiV1Routes(e, postHandlers, commentHandlers)

	// Start server
	serverUrl := ":" + appConfig.AppConfig.Port
	e.Logger.Fatal(e.Start(serverUrl))
}
