package handlers

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/domain/entity"
	"go-cource-api/interfaces/dto"
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
		return err
	}

	user := &entity.User{
		Name: userDto.Name,
		Email: userDto.Email,
	}

	_, err := u.app.Save(user)
	if err != nil {
		println("qwe")
	}

	return c.JSON(http.StatusCreated, user)
}
