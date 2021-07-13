package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	config2 "go-cource-api/application/config"
	handlers2 "go-cource-api/application/handlers"
	middlewares2 "go-cource-api/application/middlewares"
	"go-cource-api/infrustructure/persistence"
	security2 "go-cource-api/infrustructure/security"
)

func main() {
	config := config2.NewConfig()

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

	posts := handlers2.NewPosts(services.Post)
	comments := handlers2.NewComments(services.Comment)
	security := handlers2.NewSecurity(security2.NewSecurity(services.User))

	responseResponder := config2.NewResponseResponder()

	e := echo.New()
	e.Use(middlewares2.ConfigInjectorMiddleware(config))
	e.Use(middlewares2.ResponderInjectorMiddleware(responseResponder))
	e.Validator = &config2.CustomValidator{
		Validator: validator.New(),
	}
	e.Renderer = config2.Renderer()
	apiV1 := e.Group("/api/v1")
	apiV1.Use(middlewares2.SecurityContextMiddleware)
	restrictedApiV1 := apiV1.Group("")
	restrictedApiV1.Use(middlewares2.AuthMiddleware())

	// Auth
	e.GET("/login", security.UiLogin)
	e.GET("/register", security.UiRegister)
	e.GET("/auth/social/:provider", security.SocialRedirect)
	e.GET("/auth/social/:provider/success", security.SocialLoginSuccess)
	apiV1.POST("/register", security.Register)
	apiV1.POST("/login", security.Login)

	// Public API
	apiV1.GET("/posts", posts.List)
	apiV1.GET("/posts/:id", posts.FindOne)
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
