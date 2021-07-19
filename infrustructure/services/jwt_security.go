package services

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"
	"go-cource-api/infrustructure/security"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
	"net/http"
	"time"
)

type JwtSecurity struct {
	userRepo repository.UserRepository
}

func NewSecurityService(userRepo repository.UserRepository) *JwtSecurity {
	return &JwtSecurity{userRepo}
}

var _ security.TokenSecurity = &JwtSecurity{}

func (s *JwtSecurity) RegisterUser(user *entity.User) error {
	hashedPassword, _ := s.HashPassword(user.Password.String)
	user.Password = null.StringFrom(string(hashedPassword))

	if _, err := s.userRepo.Save(user); err != nil {
		return err
	}

	return nil
}

func (s *JwtSecurity) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (s *JwtSecurity) LoginUser(email string, password string) (*string, error) {
	user, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	err = s.VerifyPassword(password, user.Password.String)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return s.GenerateToken(*user)
}

func (s *JwtSecurity) VerifyPassword(plain string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}

func (s *JwtSecurity) GenerateToken(user entity.User) (*string, error) {
	claims := &security.JwtCustomClaims{
		Id:    user.Id,
		Email: user.Email.String,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(""))

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return &tokenString, nil
}

func (s *JwtSecurity) FindUserByEmail(email string) (*entity.User, error) {
	return s.userRepo.FindByEmail(email)
}
