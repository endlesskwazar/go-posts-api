package middlewares

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
)

func ResponderInjectorMiddleware(responseResponder *application.ResponseResponder) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("responder", responseResponder)
			return next(c)
		}
	}
}
