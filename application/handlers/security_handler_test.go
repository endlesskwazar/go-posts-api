package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mock "go-cource-api/application/_mocks"
	"go-cource-api/application/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	securityMock := mock.NewMockSecurityAppInterface(ctrl)
	securityMock.EXPECT().RegisterUser(gomock.Any())

	securityHandlers := NewSecurity(securityMock)

	e := BuildApp()

	registerDto := &dto.RegisterUserDto{
		Name: "test",
		Email: "test@mail.com",
		Password: "supersecret",
	}

	jsonBody, _ := json.Marshal(registerDto)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec)

	if assert.NoError(t, securityHandlers.Register(context)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	securityMock := mock.NewMockSecurityAppInterface(ctrl)
	mockedToken := "2348962u3ighbj542j34l"

	loginUserDto := &dto.LoginUserDto{
		Email: "test@mail.com",
		Password: "supersecret",
	}

	securityMock.EXPECT().LoginUser(loginUserDto.Email, loginUserDto.Password).Return(&mockedToken, nil)

	securityHandlers := NewSecurity(securityMock)

	e := BuildApp()

	jsonBody, _ := json.Marshal(loginUserDto)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec)

	if assert.NoError(t, securityHandlers.Login(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUiLogin_Success(t *testing.T) {
	e := BuildApp()
	ctrl := gomock.NewController(t)
	securityMock := mock.NewMockSecurityAppInterface(ctrl)
	securityHandlers := NewSecurity(securityMock)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec)

	if assert.NoError(t, securityHandlers.UiLogin(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUiRegister_Success(t *testing.T) {
	e := BuildApp()
	ctrl := gomock.NewController(t)
	securityMock := mock.NewMockSecurityAppInterface(ctrl)
	securityHandlers := NewSecurity(securityMock)

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := BuildContext(e, req, rec)

	if assert.NoError(t, securityHandlers.UiRegister(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}