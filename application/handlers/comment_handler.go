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
	translator := c.Get("translator").(lang.Translator)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			translator.Translate("error.url.parameter", "id"))
	}

	comments, err := h.service.FindByPostId(int64(postId))

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
// @security ApiKeyAuth
// @Param dto.CommentDto body dto.CommentDto false "Comment data"
// @Success 201 {object} entity.Comment
// @Router /api/v1/posts/{postId}/comments [post]
func (h *CommentHandlers) Save(c echo.Context) error {
	postId, err := strconv.Atoi(c.Param("postId"))
	translator := c.Get("translator").(lang.Translator)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			translator.Translate("error.url.parameter", "postId"))
	}

	commentDto := new(dto.CommentDto)
	if err := c.Bind(commentDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(commentDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	securityContext := c.(*application.SecurityContext)

	comment := &entity.Comment{
		Body:   null.StringFrom(commentDto.Body),
		PostId: null.IntFrom(int64(postId)),
		UserId: null.IntFrom(int64(securityContext.UserClaims().Id)),
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
// @security ApiKeyAuth
// @Param id path int true "Comment id"
// @Param dto.UpdateCommentDto body dto.UpdateCommentDto false "Comment data"
// @Success 200 {array} entity.Comment
// @Router /api/v1/comments/{id} [put]
func (h *CommentHandlers) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	translator := c.Get("translator").(lang.Translator)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			translator.Translate("error.url.parameter", "id"))
	}

	commentDto := new(dto.CommentDto)

	if err := c.Bind(commentDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(commentDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	comment, err := h.service.FindById(int64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	securityContext := c.(*application.SecurityContext)

	if comment.UserId.Int64 != securityContext.UserClaims().Id {
		return echo.NewHTTPError(http.StatusNotFound, "Not found")
	}

	comment.Body = null.StringFrom(commentDto.Body)

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
// @security ApiKeyAuth
// @Param id path int true "Comment id"
// @Success 204
// @Router /api/v1/comments/{id} [delete]
func (h *CommentHandlers) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	translator := c.Get("translator").(lang.Translator)

	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			translator.Translate("error.url.parameter", "id"))
	}

	securityContext := c.(*application.SecurityContext)

	comment, err := h.service.FindById(int64(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if securityContext.UserClaims().Id != comment.UserId.Int64 {
		return echo.NewHTTPError(http.StatusForbidden, "You can't perform operation")
	}

	if err = h.service.Delete(int64(id)); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
