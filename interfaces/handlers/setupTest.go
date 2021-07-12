package handlers

import (
	unsecureJWT "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/domain"
	"go-cource-api/interfaces/validation"
	"net/http"
)

func BuildApp(withValidator bool) *echo.Echo {
	app := echo.New()

	if withValidator {
		app.Validator = &validation.CustomValidator{
			Validator: validator.New(),
		}
	}

	return app
}

func BuildContext(app *echo.Echo, r *http.Request, w http.ResponseWriter, withUser bool) echo.Context {
	context := app.NewContext(r, w)

	if withUser {
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

		return &cc
	}

	return context
}
