package handlers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/domain/entity"
	"go-cource-api/interfaces/dto"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Users struct {
	app application.UserAppInterface
}

func NewUsers(app application.UserAppInterface) *Users {
	return &Users{
		app: app,
	}
}

func (u *Users) Register(c echo.Context) error  {
	println("Executing register handler")
	registerUserDto := new(dto.RegisterUserDto)

	if err := c.Bind(registerUserDto); err != nil {
		println("bind error")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(registerUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err := u.app.FindByEmail(registerUserDto.Email)

	if err != nil {
		println("find by email error")
		println(err)
	}

	// TODO: check for error
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerUserDto.Password), 14)

	user := &entity.User{
		Name: registerUserDto.Name,
		Email: registerUserDto.Email,
		Password: string(hashedPassword),
	}

	if _, err := u.app.Save(user); err != nil {
		println("error while saving user")
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.String(http.StatusOK, "qwe")
}

func (u *Users) Login(c echo.Context) error {
	loginUserDto := new(dto.LoginUserDto)

	if err := c.Bind(loginUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(loginUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := u.app.FindByEmail(loginUserDto.Email)

	// If error used with email entered doesnt exists
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUserDto.Password))

	// If error password is wrong
	if err != nil {
		return  echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.Id,
		"email": user.Email,
	})

	// Sign and get the complete encoded token as a string using the secret
	// TODO: ceck for error

	// TODO: use real cryp/rand
	var hmacSampleSecret []byte

	// Todo: check for error
	tokenString, err := token.SignedString(hmacSampleSecret)

	println(tokenString)

	type Res struct {
		Token string `json:"token"`
	}

	return c.JSON(http.StatusOK, &Res{Token: tokenString})
}

func (u *Users) List(c echo.Context) error {
	users, err := u.app.FindAll()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func (u *Users) Save(c echo.Context) error {
	userDto := new(dto.UserDto)

	if err := c.Bind(userDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(userDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := &entity.User{
		Name: userDto.Name,
		Email: userDto.Email,
		Password: userDto.Password,
	}

	_, err := u.app.Save(user)
	if err != nil {
		println("qwe")
	}

	return c.JSON(http.StatusCreated, user)
}
