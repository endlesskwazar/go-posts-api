package application

import (
	"go-cource-api/domain/entity"
	"go-cource-api/infrustructure"
)

type securityApp struct {
	security infrustructure.Security
}

var _ SecurityAppInterface = &securityApp{}

type SecurityAppInterface interface {
	HashPassword(string) ([]byte, error)
	LoginUser(string, string) (*string, error)
	VerifyPassword(string, string) error
	GenerateToken(entity.User) (*string, error)
	RegisterUser(user *entity.User) error
	IsUserExists(email string) bool
	FindUserByEmail(email string) (*entity.User, error)
}

func (app *securityApp) HashPassword(password string) ([]byte, error) {
	return app.security.HashPassword(password)
}

func (app *securityApp) LoginUser(email string, password string) (*string, error) {
	return app.security.LoginUser(email, password)
}

func (app *securityApp) VerifyPassword(plain string, hash string) error {
	return app.security.VerifyPassword(plain, hash)
}

func (app *securityApp) GenerateToken(user entity.User) (*string, error) {
	return app.security.GenerateToken(user)
}

func (app *securityApp) RegisterUser(user *entity.User) error {
	return app.security.RegisterUser(user)
}

func (app *securityApp) IsUserExists(email string) bool {
	return app.security.IsUserExists(email)
}

func (app *securityApp) FindUserByEmail(email string) (*entity.User, error) {
	return app.security.FindUserByEmail(email)
}
