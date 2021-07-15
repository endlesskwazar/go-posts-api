package security

import "go-cource-api/domain/entity"

type TokenSecurity interface {
	HashPassword(string) ([]byte, error)
	LoginUser(string, string) (*string, error)
	RegisterUser(user *entity.User) error
	VerifyPassword(string, string) error
	GenerateToken(entity.User) (*string, error)
}
