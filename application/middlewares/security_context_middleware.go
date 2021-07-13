package middlewares

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application/config"
)

func SecurityContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &config.SecurityContext{
			Context: c,
		}
		return next(cc)
	}
}
