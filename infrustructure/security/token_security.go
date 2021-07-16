package security

import "go-cource-api/domain/entity"

type TokenSecurity interface {
	HashPassword(plainPassword string) ([]byte, error)
	LoginUser(email string, password string) (*string, error)
	RegisterUser(user *entity.User) error
	VerifyPassword(plain string, hash string) error
	GenerateToken(user entity.User) (*string, error)
	FindUserByEmail(email string) (*entity.User, error)
}
