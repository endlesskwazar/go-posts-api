package application

import (
	"errors"
	"github.com/labstack/echo/v4"
)

type Responder interface {
	Respond(c echo.Context, code int, content interface{}) error
}

type ResponseResponder struct {}

var _ Responder = &ResponseResponder{}

func NewResponseResponder() *ResponseResponder {
	return &ResponseResponder{}
}

func(r *ResponseResponder) Respond(c echo.Context, code int, content interface{}) error {
	accept := c.Request().Header.Get("Accept")

	switch accept {
	case "application/json":
		return c.JSON(code, content)
	case "application/xml":
		return c.XML(code, content)
	default:
		return errors.New("unsupported accept type")
	}
}
