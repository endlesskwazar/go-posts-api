package security

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Id uint64 `json:"id"`
	Email  string `json:"name"`
	jwt.StandardClaims
}
