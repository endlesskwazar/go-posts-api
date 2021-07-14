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

type CommentHandlers struct {
	service services.CommentService
}

func NewCommentHandlers(service services.CommentService) *CommentHandlers {
	return &CommentHandlers{
		service: service,
	}
}

// FindByPostId godoc
// @Summary Returns all comments to post
// @Description Returns all comments to post
// @Tags Posts
// @Produce json,xml
// @Param postId path int true "Post id"
// @Success 200 {array} entity.Comment
// @Router /api/v1/posts/{postId}/comments [get]
func (h *CommentHandlers) FindByPostId(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("postId"))

	if err != nil {
		return err
	}

	comments, err := h.service.FindByPostId(uint64(postId))

	if err != nil {
		return err
	}

	responder := c.Get("responder").(application.Responder)
	return responder.Respond(c, http.StatusOK, comments)
}

// Save godoc
// @Summary Creates comment for post
// @Description Creates comment for post
// @Tags Posts
// @Accept xml,json
// @Produce  xml,json
// @Param dto.CommentDto body dto.CommentDto false "Comment data"
// @Success 201 {object} entity.Comment
// @Router /api/v1/posts/{postId}/comments [post]
func (h *CommentHandlers) Save(c echo.Context) error {
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

	_, err = h.service.Save(comment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	responder := c.Get("responder").(application.Responder)
	return responder.Respond(c, http.StatusCreated, comment)
}

// Update godoc
// @Summary Update comment
// @Description Update comment
// @Tags Comments
// @Accept xml,json
// @Produce json,xml
// @Param id path int true "Comment id"
// @Param dto.UpdateCommentDto body dto.UpdateCommentDto false "Comment data"
// @Success 200 {array} entity.Comment
// @Router /api/v1/comments/{id} [put]
func (h *CommentHandlers) Update(c echo.Context) error {
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

	comment, err := h.service.FindById(uint64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	securityContext := c.(*application.SecurityContext)

	if comment.UserId != securityContext.UserClaims().Id {
		return echo.NewHTTPError(http.StatusNotFound, "Not found")
	}

	comment.Body = commentDto.Body

	if err = h.service.Update(comment); err != nil {
		return err
	}

	responder := c.Get("responder").(application.Responder)
	return responder.Respond(c, http.StatusOK, comment)
}

// Delete godoc
// @Summary Delete comment
// @Description Delete comment
// @Tags Comments
// @Produce json,xml
// @Param id path int true "Comment id"
// @Success 204
// @Router /api/v1/comments/{id} [delete]
func (h *CommentHandlers) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	securityContext := c.(*application.SecurityContext)

	if err != nil {
		return err
	}

	comment, err := h.service.FindById(uint64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if securityContext.UserClaims().Id != comment.UserId {
		return echo.NewHTTPError(http.StatusForbidden, "You can't perform operation")
	}

	if err = h.service.Delete(uint64(id)); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
