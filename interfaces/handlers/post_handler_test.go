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

func TestListPost_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	postRepoMock := mock.NewMockPostRepository(ctrl)
	postRepoMock.EXPECT().FindAll()
	postHandlers := NewPosts(postRepoMock)

	e := BuildApp()

	reqJson := httptest.NewRequest(http.MethodGet, "/", nil)
	reqJson.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := BuildContext(e, reqJson, rec)

	if assert.NoError(t, postHandlers.List(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(
			t,
			echo.MIMEApplicationJSONCharsetUTF8,
			rec.Header().Get(echo.HeaderContentType),
		)
	}
}

func TestSavePost_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	postRepo := mock.NewMockPostRepository(ctrl)
	postRepo.EXPECT().Save(gomock.Any())
	postHandlers := NewPosts(postRepo)

	postDto := &dto.PostDto{
		Title: "test",
		Body:  "test",
	}
	jsonBody, _ := json.Marshal(postDto)

	e := BuildApp()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := BuildContext(e, req, rec)

	if assert.NoError(t, postHandlers.Save(c)) {
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
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := BuildContext(e, req, rec)

	c.SetPath("/api/v1/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, postHandlers.Delete(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestUpdatePost_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	postRepo := mock.NewMockPostRepository(ctrl)
	postHandlers := NewPosts(postRepo)
	standardUserId := uint64(1)
	postIdToUpdate := uint64(1)

	postRepo.
		EXPECT().
		FindByIdAndUserId(standardUserId, postIdToUpdate)

	postRepo.
		EXPECT().
		Update(gomock.Any())

	e := BuildApp()

	postDto := &dto.PostDto{
		Title: "test",
		Body:  "test",
	}

	jsonBody, _ := json.Marshal(postDto)

	req := httptest.NewRequest(http.MethodPut, "/", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec)
	context.SetPath("api/v1/posts/:id")
	context.SetParamNames("id")
	context.SetParamValues("1")

	if assert.NoError(t, postHandlers.Update(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
