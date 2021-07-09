package application

import (
	unsecureJWT "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go-cource-api/domain"
)

type SecurityContext struct {
	echo.Context
}

func (c *SecurityContext) UserClaims() *domain.JwtCustomClaims {
	user := c.Get("user").(*unsecureJWT.Token)
	claims := user.Claims.(*domain.JwtCustomClaims)

	return claims
}

