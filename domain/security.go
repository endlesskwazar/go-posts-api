package domain

import "go-cource-api/domain/entity"

type Security interface {
	HashPassword(string) ([]byte, error)
	LoginUser(string, string) (*string, error)
	RegisterUser(user *entity.User) error
	IsUserExists(email string) bool
	VerifyPassword(string, string) error
	GenerateToken(entity.User) (*string, error)
}

