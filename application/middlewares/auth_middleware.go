package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-cource-api/infrustructure/security"
)

func AuthMiddleware() echo.MiddlewareFunc {
	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte(""),
		Claims: &security.JwtCustomClaims{},
	}

	return middleware.JWTWithConfig(jwtConfig)
}
