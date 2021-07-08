package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-cource-api/domain"
)

func AuthMiddleware() echo.MiddlewareFunc {
	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte(""),
		Claims: &domain.JwtCustomClaims{},
	}

	return middleware.JWTWithConfig(jwtConfig)
}
