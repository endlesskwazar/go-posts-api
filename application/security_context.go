package application

import (
	unsecureJWT "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go-cource-api/infrustructure/security"
)

type SecurityContext struct {
	echo.Context
}

func (c *SecurityContext) UserClaims() *security.JwtCustomClaims {
	user := c.Get("user").(*unsecureJWT.Token)
	claims := user.Claims.(*security.JwtCustomClaims)

	return claims
}

