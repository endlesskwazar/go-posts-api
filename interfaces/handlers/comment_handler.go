package handlers

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/domain/entity"
	"go-cource-api/interfaces/dto"
	"net/http"
	"strconv"
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(commentDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	securityContext := c.(*application.SecurityContext)
	// TODO: handle error
	postId, _ := strconv.Atoi(c.Param("postId"))

	comment := &entity.Comment{
		Body: commentDto.Body,
		PostId: uint64(postId),
		UserId: securityContext.UserClaims().Id,
	}

	_, err := p.app.Save(comment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, comment)
}

func (p *Comments) FindByPostId(c echo.Context) error {
	// TODO: handle error
	postId, _ := strconv.Atoi(c.Param("postId"))

	comments, err := p.app.FindByPostId(uint64(postId))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comments)
}
