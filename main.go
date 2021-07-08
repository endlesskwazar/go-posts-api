package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go-cource-api/infrustructure"
	"go-cource-api/infrustructure/persistence"
	"go-cource-api/infrustructure/validation"
	"go-cource-api/interfaces/handlers"
	"go-cource-api/interfaces/middlewares"
	http "net/http"
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
	users := handlers.NewUsers(services.User)
	security := handlers.NewSecurity(infrustructure.NewSecurity(services.User))

	e := echo.New()
	e.Validator = &validation.CustomValidator{
		Validator: validator.New(),
	}

	e.Renderer = Renderer()

	e.POST("/register", security.Register)
	e.POST("/login", security.Login)

	e.POST("/api/posts", posts.Save)
	e.GET("/api/posts", posts.List)

	e.GET("/api/users", users.List)
	e.POST("/api/users", users.Save)

	r := e.Group("/api")
	r.Use(middlewares.AuthMiddleware())
	r.GET("/restricted", func(context echo.Context) error {
		return context.String(200, "qweqwe")
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{})
	})

	e.GET("/register", func(c echo.Context) error {
		return c.Render(http.StatusOK, "register.html", map[string]interface{}{})
	})

	e.GET("/auth/social/:provider", security.SocialRedirect)
	e.GET("/auth/social/:provider/success", security.SocialLoginSuccess)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
