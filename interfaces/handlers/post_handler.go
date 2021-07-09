package handlers

import (
	unsecureJWT "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/domain"
	"go-cource-api/domain/entity"
	"go-cource-api/interfaces/dto"
	"net/http"
)

type Posts struct {
	app application.PostAppInterface
}

func NewPosts(app application.PostAppInterface) *Posts {
	return &Posts{
		app: app,
	}
}

func (p *Posts) List(c echo.Context) error {
	posts, err := p.app.FindAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, posts)
}

func (p *Posts) Save(c echo.Context) error {
	postDto := new(dto.PostDto)

	if err := c.Bind(postDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(postDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*unsecureJWT.Token)
	println("Here the user from context")
	println(user)
	claims := user.Claims.(*domain.JwtCustomClaims)

	post := &entity.Post{
		Title: postDto.Title,
		Body: postDto.Body,
		UserId: claims.Id,
	}

	_, err := p.app.Save(post)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, post)
}
