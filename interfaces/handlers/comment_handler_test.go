package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-cource-api/domain/entity"
	mock "go-cource-api/interfaces/_mocks"
	"go-cource-api/interfaces/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	commentRepo := mock.NewMockCommentRepository(ctrl)
	commentHandlers := NewComments(commentRepo)

	// Todo: check for normal data
	commentRepo.
		EXPECT().
		Save(gomock.Any())

	e := BuildApp(true)

	postDto := &dto.CommentDto{
		Body: "test",
	}

	jsonBody, _ := json.Marshal(postDto)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec, true)
	context.SetParamNames("postId")
	context.SetParamValues("1")

	h := commentHandlers

	if assert.NoError(t, h.Save(context)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestFindByPostId_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	commentRepo := mock.NewMockCommentRepository(ctrl)
	commentHandlers := NewComments(commentRepo)
	postId := uint64(1)

	commentRepo.
		EXPECT().
		FindByPostId(postId)

	e := BuildApp(false)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec, false)
	context.SetParamNames("postId")
	context.SetParamValues("1")

	if assert.NoError(t, commentHandlers.FindByPostId(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	commentRepo := mock.NewMockCommentRepository(ctrl)
	commentHandlers := NewComments(commentRepo)

	commentRepo.EXPECT().FindById(uint64(1))
	commentRepo.EXPECT().Delete(uint64(1))

	e := BuildApp(true)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec, true)
	context.SetParamNames("id")
	context.SetParamValues("1")

	h := commentHandlers

	if assert.NoError(t, h.Delete(context)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestUpdateComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	commentRepo := mock.NewMockCommentRepository(ctrl)

	commentRepo.EXPECT().FindById(uint64(1)).Return(&entity.Comment{UserId: uint64(1)}, nil)
	commentRepo.EXPECT().Update(gomock.Any())

	commentHandlers := NewComments(commentRepo)

	e := BuildApp(true)

	updateCommentDto := &dto.UpdateCommentDto{
		Body: "test",
	}

	jsonBody, _ := json.Marshal(updateCommentDto)

	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec, true)
	context.SetParamNames("id")
	context.SetParamValues("1")

	h := commentHandlers

	if assert.NoError(t, h.Update(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}