package handlers

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application"
	"go-cource-api/application/dto"
	"go-cource-api/application/services"
	"go-cource-api/domain/entity"
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

	if err != nil {
		return err
	}

	post, err := h.service.FindById(uint64(id))

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
		Title: postDto.Title,
		Body: postDto.Body,
		UserId: securityContext.UserClaims().Id,
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
// @Param id path int true "Post id"
// @Success 204
// @Router /api/v1/posts/{id} [delete]
func (h *PostHandlers) Delete(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	securityContext := c.(*application.SecurityContext)
	userId := securityContext.UserClaims().Id

	_, err = h.service.FindByIdAndUserId(uint64(postId), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	if err = h.service.Delete(uint64(postId)); err != nil {
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
// @Param dto.UpdatePostDto body dto.UpdatePostDto false "Post data"
// @Param id path int true "Post id"
// @Success 200 {object} entity.Post
// @Router /api/v1/posts/{id} [put]
func (h *PostHandlers) Update(c echo.Context) error {
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

	_, err := h.service.FindByIdAndUserId(uint64(postId), userId)

	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, err)
	}

	updatedPost := &entity.Post{
		Id: uint64(postId),
		Title: postDto.Title,
		Body: postDto.Body,
	}

	if err = h.service.Update(updatedPost); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	responder := c.Get("responder").(application.Responder)
	return responder.Respond(c, http.StatusOK, updatedPost)
}
