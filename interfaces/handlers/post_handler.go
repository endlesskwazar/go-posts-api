package handlers

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
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
		return err
	}

	return c.JSON(http.StatusOK, posts)
}

func (p *Posts) Save(c echo.Context) error {
	postDto := new(dto.PostDto)

	if err := c.Bind(postDto); err != nil {
		return err
	}

	if err := c.Validate(postDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	post := &entity.Post{
		Title: postDto.Title,
		UserId: postDto.UserId,
	}

	_, err := p.app.Save(post)
	if err != nil {
		println("qwe")
	}

	return c.JSON(http.StatusCreated, post)
}
