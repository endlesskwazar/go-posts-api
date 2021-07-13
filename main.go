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
	apiV1 := e.Group("/api/v1")
	apiV1.Use(middlewares.SecurityContextMiddleware)
	restrictedApiV1 := apiV1.Group("")
	restrictedApiV1.Use(middlewares.AuthMiddleware())

	// Auth
	e.GET("/login", securityHandlers.UiLogin)
	e.GET("/register", securityHandlers.UiRegister)
	e.GET("/auth/social/:provider", securityHandlers.SocialRedirect)
	e.GET("/auth/social/:provider/success", securityHandlers.SocialLoginSuccess)
	apiV1.POST("/register", securityHandlers.Register)
	apiV1.POST("/login", securityHandlers.Login)

	// Public API
	apiV1.GET("/posts", postHandlers.List)
	apiV1.GET("/posts/:id", postHandlers.FindOne)
	apiV1.GET("/posts/:postId/comments", commentHandlers.FindByPostId)

	// Private API
	restrictedApiV1.POST("/posts", postHandlers.Save)
	restrictedApiV1.DELETE("/posts/:id", postHandlers.Delete)
	restrictedApiV1.PUT("/posts/:id", postHandlers.Update)

	restrictedApiV1.POST("/posts/:postId/comments", commentHandlers.Save)
	restrictedApiV1.DELETE("/comments/:id", commentHandlers.Delete)
	restrictedApiV1.PUT("/comments/:id", commentHandlers.Update)

	// Start server
	serverUrl := ":" + appConfig.AppConfig.Port
	e.Logger.Fatal(e.Start(serverUrl))
}
