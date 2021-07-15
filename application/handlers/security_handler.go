package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"go-cource-api/application/config"
	"go-cource-api/application/dto"
	"go-cource-api/domain/entity"
	"go-cource-api/infrustructure/security"
	"golang.org/x/oauth2"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
)

type SecurityHandlers struct {
	service security.TokenSecurity
}

func NewSecurityHandlers(service security.TokenSecurity) *SecurityHandlers {
	return &SecurityHandlers{
		service: service,
	}
}


// Register godoc
// @Summary Register new user
// @Description Register new user
// @Tags Auth
// @Accept xml,json
// @Param dto.RegisterUserDto body dto.RegisterUserDto false "Register data"
// @Success 204
// @Router /register [post]
func (h *SecurityHandlers) Register(c echo.Context) error {
	registerUserDto := new(dto.RegisterUserDto)

	if err := c.Bind(registerUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(registerUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := &entity.User{
		Name:     null.StringFrom(registerUserDto.Name),
		Email:    null.StringFrom(registerUserDto.Email),
		Password: null.StringFrom(registerUserDto.Password),
	}

	err := h.service.RegisterUser(user)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// Login godoc
// @Summary Login user
// @Description Login user
// @Tags Auth
// @Accept xml,json
// @Produce xml,json
// @Param dto.LoginUserDto body dto.LoginUserDto false "Register data"
// @Success 200 {object} security.Token
// @Router /login [post]
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


	return echo.NewHTTPError(http.StatusInternalServerError, "Social auth error")
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
			Email:    null.StringFrom(result["email"].(string)),
			Name:     null.StringFrom(result["name"].(string)),
			Password: null.StringFrom(randstr.Hex(16)),
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
