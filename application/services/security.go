package services

import (
	"go-cource-api/domain/entity"
	"go-cource-api/infrustructure/security"
)

type securityService struct {
	security security.JwtSecurity
}

var _ SecurityService = &securityService{}

type SecurityService interface {
	HashPassword(string) ([]byte, error)
	LoginUser(string, string) (*string, error)
	VerifyPassword(string, string) error
	GenerateToken(entity.User) (*string, error)
	RegisterUser(user *entity.User) error
	IsUserExists(email string) bool
	FindUserByEmail(email string) (*entity.User, error)
}

func (app *securityService) HashPassword(password string) ([]byte, error) {
	return app.security.HashPassword(password)
}

func (app *securityService) LoginUser(email string, password string) (*string, error) {
	return app.security.LoginUser(email, password)
}

func (app *securityService) VerifyPassword(plain string, hash string) error {
	return app.security.VerifyPassword(plain, hash)
}

func (app *securityService) GenerateToken(user entity.User) (*string, error) {
	return app.security.GenerateToken(user)
}

func (app *securityService) RegisterUser(user *entity.User) error {
	return app.security.RegisterUser(user)
}

func (app *securityService) IsUserExists(email string) bool {
	return app.security.IsUserExists(email)
}

func (app *securityService) FindUserByEmail(email string) (*entity.User, error) {
	return app.security.FindUserByEmail(email)
}
