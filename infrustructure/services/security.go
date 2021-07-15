package services

import (
	"go-cource-api/domain/entity"
	"go-cource-api/infrustructure/security"
)

type securityService struct {
	strategy security.JwtSecurity
}

var _ SecurityService = &securityService{}

type SecurityService interface {
	HashPassword(string) ([]byte, error)
	LoginUser(string, string) (*string, error)
	VerifyPassword(string, string) error
	GenerateToken(entity.User) (*string, error)
	RegisterUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
}

func (s *securityService) HashPassword(password string) ([]byte, error) {
	return s.strategy.HashPassword(password)
}

func (s *securityService) LoginUser(email string, password string) (*string, error) {
	return s.strategy.LoginUser(email, password)
}

func (s *securityService) VerifyPassword(plain string, hash string) error {
	return s.strategy.VerifyPassword(plain, hash)
}

func (s *securityService) GenerateToken(user entity.User) (*string, error) {
	return s.strategy.GenerateToken(user)
}

func (s *securityService) RegisterUser(user *entity.User) error {
	return s.strategy.RegisterUser(user)
}

func (s *securityService) FindUserByEmail(email string) (*entity.User, error) {
	return s.strategy.FindUserByEmail(email)
}
