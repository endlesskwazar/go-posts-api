package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go-cource-api/infrustructure"
	"go-cource-api/infrustructure/persistence"
	"go-cource-api/infrustructure/validation"
	"go-cource-api/interfaces/handlers"
	"go-cource-api/interfaces/middlewares"
	"os"
)

func main() {
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	services, err := persistence.NewRepositories(user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	services.Automigrate()

	posts := handlers.NewPosts(services.Post)
	security := handlers.NewSecurity(infrustructure.NewSecurity(services.User))

	e := echo.New()
	e.Validator = &validation.CustomValidator{
		Validator: validator.New(),
	}
	e.Renderer = Renderer()
	apiV1 := e.Group("/api/v1")
	restrictedApiV1 := apiV1.Group("")
	restrictedApiV1.Use(middlewares.AuthMiddleware())

	// Auth
	e.GET("/login", security.UiLogin)
	e.GET("/register", security.UiRegister)
	apiV1.POST("/register", security.Register)
	apiV1.POST("/login", security.Login)
	e.GET("/auth/social/:provider", security.SocialRedirect)
	e.GET("/auth/social/:provider/success", security.SocialLoginSuccess)

	// Public API
	e.GET("/api/posts", posts.List)

	// Private API
	e.POST("/api/posts", posts.Save)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
