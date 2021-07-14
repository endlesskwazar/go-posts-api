package security

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Id int64 `json:"id"`
	Email  string `json:"name"`
	jwt.StandardClaims
}
