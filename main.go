package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/application/config"
	"go-cource-api/application/handlers"
	"go-cource-api/application/middlewares"
	_ "go-cource-api/docs"
	"go-cource-api/infrustructure/persistence"
	"go-cource-api/infrustructure/services"
	"go-cource-api/routes"
)

var (
	appConfig *config.Config
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
	appConfig = config.NewConfig()

	repositories, err := buildRepositories()
	if err != nil {
		panic(err)
	}

	postHandlers := handlers.NewPostHandlers(repositories.Post)
	commentHandlers := handlers.NewCommentHandlers(repositories.Comment)
	securityHandlers := handlers.NewSecurityHandlers(services.NewSecurityService(repositories.User))

	e := buildApp()

	routes.InitAuthRoutes(e, securityHandlers)
	routes.InitApiV1Routes(e, postHandlers, commentHandlers)
	routes.InitSwaggerRoute(e)

	e.Logger.Fatal(e.Start(appConfig.AppConfig.Address))
}

func buildRepositories() (*persistence.Repositories, error) {
	repositories, err := persistence.NewRepositories(appConfig.DatabaseConfig)

	if err != nil {
		return nil, err
	}

	if err = repositories.Automigrate(); err != nil {
		return nil, err
	}

	return repositories, nil
}

func buildApp() *echo.Echo {
	e := echo.New()
	responder := application.NewResponseResponder()

	e.Use(middlewares.ConfigInjectorMiddleware(appConfig))
	e.Use(middlewares.ResponderInjectorMiddleware(responder))
	e.Validator = &application.CustomValidator{
		Validator: validator.New(),
	}
	e.HTTPErrorHandler = application.ErrorHandler
	e.Renderer = application.Renderer()

	return e
}
