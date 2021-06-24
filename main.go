package main

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/infrustructure/persistence"
	"go-cource-api/interfaces/handlers"
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

	e := echo.New()

	e.POST("/api/posts", posts.Save)
	e.GET("/api/posts", posts.List)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
