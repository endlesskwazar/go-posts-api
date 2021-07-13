package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mock "go-cource-api/application/_mocks"
	"go-cource-api/application/dto"
	"go-cource-api/domain/entity"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindByPostId_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	commentRepo := mock.NewMockCommentRepository(ctrl)
	commentHandlers := NewCommentHandlers(commentRepo)
	postId := uint64(1)

	commentRepo.
		EXPECT().
		FindByPostId(postId)

	e := BuildApp()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec)
	context.SetParamNames("postId")
	context.SetParamValues("1")

	if assert.NoError(t, commentHandlers.FindByPostId(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get(echo.HeaderContentType))
	}
}

func TestSaveComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	commentRepo := mock.NewMockCommentRepository(ctrl)
	commentHandlers := NewCommentHandlers(commentRepo)

	commentRepo.EXPECT().Save(gomock.Any())

	e := BuildApp()

	commentDto := &dto.CommentDto{
		Body: "test",
	}

	jsonBody, _ := json.Marshal(commentDto)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec)
	context.SetParamNames("postId")
	context.SetParamValues("1")

	if assert.NoError(t, commentHandlers.Save(context)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get(echo.HeaderContentType))
	}
}

func TestUpdateComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	commentRepo := mock.NewMockCommentRepository(ctrl)

	commentRepo.EXPECT().FindById(uint64(1)).Return(&entity.Comment{UserId: uint64(1)}, nil)
	commentRepo.EXPECT().Update(gomock.Any())

	commentHandlers := NewCommentHandlers(commentRepo)

	e := BuildApp()

	updateCommentDto := &dto.UpdateCommentDto{
		Body: "test",
	}

	jsonBody, _ := json.Marshal(updateCommentDto)
	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec)
	context.SetParamNames("id")
	context.SetParamValues("1")

	h := commentHandlers

	if assert.NoError(t, h.Update(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, echo.MIMEApplicationJSONCharsetUTF8, rec.Header().Get(echo.HeaderContentType))
	}
}

func TestDeleteComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	commentRepo := mock.NewMockCommentRepository(ctrl)
	commentHandlers := NewCommentHandlers(commentRepo)

	commentRepo.EXPECT().FindById(uint64(1)).Return(&entity.Comment{UserId: uint64(1)}, nil)
	commentRepo.EXPECT().Delete(uint64(1))

	e := BuildApp()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := BuildContext(e, req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, commentHandlers.Delete(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}