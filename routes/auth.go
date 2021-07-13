package routes

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application/handlers"
)

func InitAuthRoutes(e *echo.Echo, handlers *handlers.SecurityHandlers) {
	e.GET("/login", handlers.UiLogin)
	e.GET("/register", handlers.UiRegister)
	e.GET("/auth/social/:provider", handlers.SocialRedirect)
	e.GET("/auth/social/:provider/success", handlers.SocialLoginSuccess)
	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)
}
