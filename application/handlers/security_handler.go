package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"go-cource-api/application/config"
	"go-cource-api/application/dto"
	"go-cource-api/application/services"
	"go-cource-api/domain/entity"
	"go-cource-api/infrustructure/security"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
)

type SecurityHandlers struct {
	service services.SecurityService
}

func NewSecurity(service services.SecurityService) *SecurityHandlers {
	return &SecurityHandlers{
		service: service,
	}
}

func (h *SecurityHandlers) Register(c echo.Context) error {
	registerUserDto := new(dto.RegisterUserDto)

	if err := c.Bind(registerUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(registerUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := &entity.User{
		Name:     registerUserDto.Name,
		Email:    registerUserDto.Email,
		Password: registerUserDto.Password,
	}

	err := h.service.RegisterUser(user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (h *SecurityHandlers) Login(c echo.Context) error {
	loginUserDto := new(dto.LoginUserDto)

	if err := c.Bind(loginUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(loginUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tokenString, err := h.service.LoginUser(loginUserDto.Email, loginUserDto.Password)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &security.Token{Token: *tokenString})
}

func (h *SecurityHandlers) SocialRedirect(c echo.Context) error {
	appConfig := c.Get("config").(*config.Config)
	var redirectUrl string

	switch provider := c.Param("provider"); provider {
	case "google":
		redirectUrl = appConfig.GoogleOauthConfig.AuthCodeURL("state")
	case "facebook":
		redirectUrl = appConfig.FaceBookOauthConfig.AuthCodeURL("state")
	}

	if err := c.Redirect(http.StatusFound, redirectUrl); err != nil {
		return err
	}

	// TODO: return error
	return c.String(200, "qweqwe")
}

func (h *SecurityHandlers) SocialLoginSuccess(c echo.Context) error {
	provider := c.Param("provider")
	code := c.QueryParam("code")
	appConfig := c.Get("config").(*config.Config)

	var userInfoUrl string

	switch provider {
	case "google":
		exchange, err := appConfig.GoogleOauthConfig.Exchange(oauth2.NoContext, code)
		if err != nil {
			return fmt.Errorf("code exchange failed: %s", err.Error())
		}
		userInfoUrl = "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + exchange.AccessToken
	case "facebook":
		exchange, err := appConfig.FaceBookOauthConfig.Exchange(oauth2.NoContext, code)
		if err != nil {
			return fmt.Errorf("code exchange failed: %s", err.Error())
		}
		userInfoUrl = "https://graph.facebook.com/me?fields=name,first_name,last_name,email&access_token=" + exchange.AccessToken
	}

	userInfo, err := getUserInfo(userInfoUrl)

	if err != nil {
		return err
	}

	var result map[string]interface{}
	err = json.Unmarshal(userInfo, &result)

	if err != nil {
		return err
	}

	user, err := h.service.FindUserByEmail(result["email"].(string))

	if err != nil {
		println(err.Error())
	}

	// No user in Db -> create
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		user = &entity.User{
			Email:    result["email"].(string),
			Name:     result["name"].(string),
			Password: randstr.Hex(16),
		}

		err := h.service.RegisterUser(user)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
	}

	token, err := h.service.GenerateToken(*user)

	if err != nil {
		return err
	}

	type Res struct {
		Token string `json:"token"`
	}

	return c.JSON(200, &Res{Token: *token})
}

func getUserInfo(userDetailsUrl string) ([]byte, error) {
	response, err := http.Get(userDetailsUrl)

	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}

func (h *SecurityHandlers) UiLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{})
}

func (h *SecurityHandlers) UiRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", map[string]interface{}{})
}
