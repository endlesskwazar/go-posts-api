package handlers

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application/config"
	dto2 "go-cource-api/application/dto"
	"go-cource-api/application/services"
	"go-cource-api/domain/entity"
	"net/http"
	"strconv"
)

type Comments struct {
	app services.CommentAppInterface
}

func NewComments(app services.CommentAppInterface) *Comments {
	return &Comments{
		app: app,
	}
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

	responder := c.Get("responseResponder").(config.Responder)
	return responder.Respond(c, http.StatusOK, comments)
}

func (p *Comments) Save(c echo.Context) error {
	commentDto := new(dto2.CommentDto)

	if err := c.Bind(commentDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(commentDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	securityContext := c.(*config.SecurityContext)

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

	responder := c.Get("responseResponder").(config.Responder)
	return responder.Respond(c, http.StatusCreated, comment)
}

func (p *Comments) Update(c echo.Context) error {
	commentDto := new(dto2.CommentDto)

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

	securityContext := c.(*config.SecurityContext)

	if comment.UserId != securityContext.UserClaims().Id {
		return echo.NewHTTPError(http.StatusNotFound, "Not found")
	}

	comment.Body = commentDto.Body

	if err = p.app.Update(comment); err != nil {
		return err
	}

	responder := c.Get("responseResponder").(config.Responder)
	return responder.Respond(c, http.StatusOK, comment)
}

func (p *Comments) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	cc := c.(*config.SecurityContext)

	if err != nil {
		return err
	}

	comment, err := p.app.FindById(uint64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if cc.UserClaims().Id != comment.UserId {
		return echo.NewHTTPError(http.StatusForbidden, "You can't perform operation")
	}

	if err = p.app.Delete(uint64(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusNoContent)
}
