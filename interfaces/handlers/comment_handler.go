package handlers

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/domain/entity"
	"go-cource-api/interfaces/dto"
	"net/http"
)

type Comments struct {
	app application.CommentAppInterface
}

func NewComments(app application.CommentAppInterface) *Comments {
	return &Comments{
		app: app,
	}
}

func (p *Comments) List(c echo.Context) error {
	comments, err := p.app.FindAll()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comments)
}

func (p *Comments) Save(c echo.Context) error {
	commentDto := new(dto.CommentDto)

	if err := c.Bind(commentDto); err != nil {
		return err
	}

	comment := &entity.Comment{
		UserId: commentDto.UserId,
		PostId: commentDto.PostId,
		Body: commentDto.Body,
	}

	_, err := p.app.Save(comment)
	if err != nil {
		println("qwe")
	}

	return c.JSON(http.StatusCreated, comment)
}
