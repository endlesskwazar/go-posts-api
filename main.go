package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/infrustructure"
	"go-cource-api/infrustructure/persistence"
	"go-cource-api/interfaces/handlers"
	"go-cource-api/interfaces/middlewares"
	"go-cource-api/interfaces/validation"
)

func main() {
	config := application.NewConfig()

	services, err := persistence.NewRepositories(
		config.DatabaseConfig.User,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Port,
		config.DatabaseConfig.Host,
		config.DatabaseConfig.DbName,
	)

	if err != nil {
		panic(err)
	}

	if err = services.Automigrate(); err != nil {
		panic(err)
	}

	posts := handlers.NewPosts(services.Post)
	comments := handlers.NewComments(services.Comment)
	security := handlers.NewSecurity(infrustructure.NewSecurity(services.User))

	e := echo.New()
	e.Use(middlewares.ConfigInjectorMiddleware(config))
	e.Validator = &validation.CustomValidator{
		Validator: validator.New(),
	}
	e.Renderer = Renderer()
	apiV1 := e.Group("/api/v1")
	apiV1.Use(middlewares.SecurityContextMiddleware)
	restrictedApiV1 := apiV1.Group("")
	restrictedApiV1.Use(middlewares.AuthMiddleware())

	// Auth
	e.GET("/login", security.UiLogin)
	e.GET("/register", security.UiRegister)
	e.GET("/auth/social/:provider", security.SocialRedirect)
	e.GET("/auth/social/:provider/success", security.SocialLoginSuccess)
	apiV1.POST("/register", security.Register)
	apiV1.POST("/login", security.Login)

	// Public API
	apiV1.GET("/posts", posts.List)
	apiV1.GET("/posts/:postId/comments", comments.FindByPostId)

	// Private API
	restrictedApiV1.POST("/posts", posts.Save)
	restrictedApiV1.DELETE("/posts/:id", posts.Delete)
	restrictedApiV1.PUT("/posts/:id", posts.Update)

	restrictedApiV1.POST("/posts/:postId/comments", comments.Save)
	restrictedApiV1.DELETE("/comments/:id", comments.Delete)
	restrictedApiV1.PUT("/comments/:id", comments.Update)

	// Start server
	serverUrl := ":" + config.AppConfig.Port
	e.Logger.Fatal(e.Start(serverUrl))
}
