package handlers

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/application/dto"
	"go-cource-api/application/lang"
	"go-cource-api/domain/entity"
	"go-cource-api/infrustructure/services"
	"gopkg.in/guregu/null.v4"
	"net/http"
	"strconv"
)

type PostHandlers struct {
	service services.PostService
}

func NewPostHandlers(service services.PostService) *PostHandlers {
	return &PostHandlers{
		service: service,
	}
}

// List godoc
// @Summary Get all posts
// @Description Get all posts
// @Tags Posts
// @Produce json,xml
// @Success 200 {array} entity.Post
// @Router /api/v1/posts [get]
func (h *PostHandlers) List(c echo.Context) error {
	posts, err := h.service.FindAll()
	responseResponder := c.Get("responder").(application.Responder)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return responseResponder.Respond(c, http.StatusOK, posts)
}

// FindOne godoc
// @Summary Get one post by id
// @Description Get one post by id
// @Tags Posts
// @Produce json,xml
// @Param id path int true "Post id"
// @Success 200 {object} entity.Post
// @Router /api/v1/posts/{id} [get]
func (h *PostHandlers) FindOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	translator := c.Get("translator").(lang.Translator)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			translator.Translate("error.url.parameter", "id"))
	}

	post, err := h.service.FindById(int64(id))

	if err != nil {
		return err
	}

	responder := c.Get("responder").(application.Responder)
	return responder.Respond(c, http.StatusOK, post)
}

// Save godoc
// @Summary Create post
// @Description Create post
// @Tags Posts
// @Accept xml,json
// @Produce  xml,json
// @security ApiKeyAuth
// @Param dto.PostDto body dto.PostDto false "Post data"
// @Success 201 {object} entity.Comment
// @Router /api/v1/posts [post]
func (h *PostHandlers) Save(c echo.Context) error {
	postDto := new(dto.PostDto)

	if err := c.Bind(postDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(postDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	securityContext := c.(*application.SecurityContext)

	post := &entity.Post{
		Title:  null.StringFrom(postDto.Title),
		Body:   null.StringFrom(postDto.Body),
		UserId: null.IntFrom(int64(securityContext.UserClaims().Id)),
	}

	_, err := h.service.Save(post)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	responder := c.Get("responder").(application.Responder)

	return responder.Respond(c, http.StatusCreated, post)
}

// Delete godoc
// @Summary Delete post
// @Description Delete post
// @Tags Posts
// @security ApiKeyAuth
// @Param id path int true "Post id"
// @Success 204
// @Router /api/v1/posts/{id} [delete]
func (h *PostHandlers) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	translator := c.Get("translator").(lang.Translator)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			translator.Translate("error.url.parameter", "id"))
	}

	securityContext := c.(*application.SecurityContext)
	userId := securityContext.UserClaims().Id

	_, err = h.service.FindByIdAndUserId(int64(id), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	if err = h.service.Delete(int64(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	responder := c.Get("responder").(application.Responder)
	return responder.Respond(c, http.StatusNoContent, nil)
}

// Update godoc
// @Summary Update post
// @Description Update post
// @Tags Posts
// @Accept xml,json
// @Produce  xml,json
// @security ApiKeyAuth
// @Param dto.UpdatePostDto body dto.UpdatePostDto false "Post data"
// @Param id path int true "Post id"
// @Success 200 {object} entity.Post
// @Router /api/v1/posts/{id} [put]
func (h *PostHandlers) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	translator := c.Get("translator").(lang.Translator)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			translator.Translate("error.url.parameter", "id"))
	}

	securityContext := c.(*application.SecurityContext)
	userId := securityContext.UserClaims().Id

	postDto := new(dto.PostDto)

	if err := c.Bind(postDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(postDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	_, err = h.service.FindByIdAndUserId(int64(id), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	updatedPost := &entity.Post{
		Id:    int64(id),
		Title: null.StringFrom(postDto.Title),
		Body:  null.StringFrom(postDto.Body),
	}

	if err = h.service.Update(updatedPost); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	responder := c.Get("responder").(application.Responder)
	return responder.Respond(c, http.StatusOK, updatedPost)
}
