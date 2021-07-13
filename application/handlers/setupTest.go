package handlers

import (
	unsecureJWT "github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go-cource-api/application/config"
	"go-cource-api/infrustructure/security"
	"net/http"
)

func BuildApp() *echo.Echo {
	app := echo.New()
	app.Validator = &config.CustomValidator{
		Validator: validator.New(),
	}
	app.Renderer = config.Renderer()

	return app
}

func BuildContext(app *echo.Echo, r *http.Request, w http.ResponseWriter) echo.Context {
	context := app.NewContext(r, w)

	cc := config.SecurityContext{
		Context: context,
	}

	claims := &security.JwtCustomClaims{
		Id: 1,
		Email: "test@mail.com",
	}

	user :=  &unsecureJWT.Token{
		Claims:claims,
	}

	context.Set("user", user)
	context.Set("responseResponder", config.NewResponseResponder())

	return &cc
}
