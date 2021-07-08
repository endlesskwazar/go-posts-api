package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-cource-api/application"
)

func AuthMiddleware() echo.MiddlewareFunc {
	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte(""),
		Claims: &application.JwtCustomClaims{},
	}

	return middleware.JWTWithConfig(jwtConfig)
}
