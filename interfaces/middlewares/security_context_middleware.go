package middlewares

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
)

func SecurityContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &application.SecurityContext{
			Context: c,
		}
		return next(cc)
	}
}
