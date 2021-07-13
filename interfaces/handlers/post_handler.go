package handlers

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/domain/entity"
	"go-cource-api/interfaces/dto"
	"net/http"
	"strconv"
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
	responseResponder := c.Get("responseResponder").(application.Responder)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return responseResponder.Respond(c, http.StatusOK, posts)
}

func (p *Posts) FindOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	post, err := p.app.FindById(uint64(id))

	if err != nil {
		return err
	}

	responder := c.Get("responseResponder").(application.Responder)
	return responder.Respond(c, http.StatusOK, post)
}

func (p *Posts) Save(c echo.Context) error {
	postDto := new(dto.PostDto)

	if err := c.Bind(postDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(postDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	securityContext := c.(*application.SecurityContext)

	post := &entity.Post{
		Title: postDto.Title,
		Body: postDto.Body,
		UserId: securityContext.UserClaims().Id,
	}

	_, err := p.app.Save(post)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	responder := c.Get("responseResponder").(application.Responder)

	return responder.Respond(c, http.StatusCreated, post)
}

func (p *Posts) Delete(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	securityContext := c.(*application.SecurityContext)
	userId := securityContext.UserClaims().Id

	_, err = p.app.FindByIdAndUserId(uint64(postId), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	if err = p.app.Delete(uint64(postId)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	responder := c.Get("responseResponder").(application.Responder)
	return responder.Respond(c, http.StatusNoContent, nil)
}

func (p *Posts) Update(c echo.Context) error {
	postId, _ := strconv.Atoi(c.Param("id"))
	securityContext := c.(*application.SecurityContext)
	userId := securityContext.UserClaims().Id

	postDto := new(dto.PostDto)

	if err := c.Bind(postDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(postDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err := p.app.FindByIdAndUserId(uint64(postId), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	updatedPost := &entity.Post{
		Id: uint64(postId),
		Title: postDto.Title,
		Body: postDto.Body,
	}

	if err = p.app.Update(updatedPost); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	responder := c.Get("responseResponder").(application.Responder)
	return responder.Respond(c, http.StatusOK, updatedPost)
}
