package middlewares

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application/config"
)

func ConfigInjectorMiddleware(config *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", config)
			return next(c)
		}
	}
}
