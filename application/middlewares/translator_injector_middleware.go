package middlewares

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application/lang"
)

func TranslatorInjectorMiddleware(translator lang.Translator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("translator", translator)
			return next(c)
		}
	}
}
