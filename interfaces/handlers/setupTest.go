package handlers

import (
	unsecureJWT "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/domain"
	"go-cource-api/interfaces"
	"go-cource-api/interfaces/validation"
	"net/http"
)

func BuildApp() *echo.Echo {
	app := echo.New()
	app.Validator = &validation.CustomValidator{
		Validator: validator.New(),
	}
	app.Renderer = interfaces.Renderer()

	return app
}

func BuildContext(app *echo.Echo, r *http.Request, w http.ResponseWriter) echo.Context {
	context := app.NewContext(r, w)

	cc := application.SecurityContext{
		Context: context,
	}

	claims := &domain.JwtCustomClaims{
		Id: 1,
		Email: "test@mail.com",
	}

	user :=  &unsecureJWT.Token{
		Claims:claims,
	}

	context.Set("user", user)
	context.Set("responseResponder", application.NewResponseResponder())

	return &cc
}
