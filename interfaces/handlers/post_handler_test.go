package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-cource-api/interfaces/_mocks"
	"go-cource-api/interfaces/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePost_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	postRepoMock := mock.NewMockPostRepository(ctrl)
	postHandlers := NewPosts(postRepoMock)

	postRepoMock.
		EXPECT().
		FindAll()

	e := BuildApp()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := postHandlers

	// Assertions
	if assert.NoError(t, h.List(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSavePost_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	postRepo := mock.NewMockPostRepository(ctrl)
	postHandlers := NewPosts(postRepo)

	postRepo.
		EXPECT().
		Save(gomock.Any())

	e := BuildApp()

	postDto := &dto.PostDto{
		Title: "test",
		Body: "test",
	}

	jsonBody, _ := json.Marshal(postDto)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec, true)

	h := postHandlers

	if assert.NoError(t, h.Save(context)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestDeletePost_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	postRepo := mock.NewMockPostRepository(ctrl)
	postHandlers := NewPosts(postRepo)
	idToDelete := uint64(1)
	standardUserId := uint64(1)

	postRepo.
		EXPECT().
		FindByIdAndUserId(idToDelete, standardUserId)

	postRepo.
		EXPECT().
		Delete(idToDelete)

	e := BuildApp()

	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec, true)
	context.SetPath("/api/v1/posts/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	h := postHandlers

	if assert.NoError(t, h.Delete(context)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestUpdatePost_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	postRepo := mock.NewMockPostRepository(ctrl)
	postHandlers := NewPosts(postRepo)
	standartUserId := uint64(1)
	postIdToUpdate := uint64(1)

	postRepo.
		EXPECT().
		FindByIdAndUserId(standartUserId, postIdToUpdate)

	postRepo.
		EXPECT().
		Update(gomock.Any())

	e := BuildApp()

	postDto := &dto.PostDto{
		Title: "test",
		Body: "test",
	}

	jsonBody, _ := json.Marshal(postDto)

	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec, true)
	context.SetPath("api/v1/posts/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	h := postHandlers

	if assert.NoError(t, h.Update(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}