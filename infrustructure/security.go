package infrustructure

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-cource-api/domain"
	"go-cource-api/domain/entity"
	"go-cource-api/domain/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Security struct {
	userRepo repository.UserRepository
}

func NewSecurity(userRepo repository.UserRepository) *Security {
	return &Security{userRepo}
}

var _ domain.Security = &Security{}

func (s *Security) IsUserExists(email string) bool {
	_, err := s.userRepo.FindByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}

	return true
}

func (s *Security) RegisterUser(user *entity.User) error {
	hashedPassword, _ := s.HashPassword(user.Password)
	user.Password = string(hashedPassword)

	if _, err := s.userRepo.Save(user); err != nil {
		return err
	}

	return nil
}

func (s *Security) HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (s *Security) LoginUser(email string, password string) (*string, error) {
	user, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	err = s.VerifyPassword(password, user.Password)

	if err != nil {
		return  nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return s.GenerateToken(*user)
}

func (s *Security) VerifyPassword(plain string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}

func (s *Security) GenerateToken(user entity.User) (*string, error) {
	claims := &domain.JwtCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(""))

	if err != nil {
		return  nil, echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return &tokenString, nil
}

func (s *Security) FindUserByEmail(email string) (*entity.User, error) {
	return s.userRepo.FindByEmail(email)
}
