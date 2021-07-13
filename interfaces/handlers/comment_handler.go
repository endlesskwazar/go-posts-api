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

func (p *Comments) Save(c echo.Context) error {
	commentDto := new(dto.CommentDto)

	if err := c.Bind(commentDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(commentDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	securityContext := c.(*application.SecurityContext)

	postId, err := strconv.Atoi(c.Param("postId"))

	if err != nil {
		return err
	}

	comment := &entity.Comment{
		Body: commentDto.Body,
		PostId: uint64(postId),
		UserId: securityContext.UserClaims().Id,
	}

	_, err = p.app.Save(comment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, comment)
}

func (p *Comments) FindByPostId(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("postId"))

	if err != nil {
		return err
	}

	comments, err := p.app.FindByPostId(uint64(postId))

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comments)
}

func (p *Comments) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	_, err = p.app.FindById(uint64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if err = p.app.Delete(uint64(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (p *Comments) Update(c echo.Context) error {
	commentDto := new(dto.CommentDto)

	if err := c.Bind(commentDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(commentDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	comment, err := p.app.FindById(uint64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	securityContext := c.(*application.SecurityContext)

	// TODO: extruct to own error
	if comment.UserId != securityContext.UserClaims().Id {
		return echo.NewHTTPError(http.StatusNotFound, "Not found")
	}

	comment.Body = commentDto.Body

	if err = p.app.Update(comment); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comment)
}
