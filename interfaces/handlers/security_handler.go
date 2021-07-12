package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/thanhpk/randstr"
	"go-cource-api/application"
	"go-cource-api/domain/entity"
	"go-cource-api/interfaces/dto"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	googleOauthConfig *oauth2.Config
	facebookOauthConfig *oauth2.Config
)

type Security struct {
	app application.SecurityAppInterface
}

func NewSecurity(app application.SecurityAppInterface) *Security {
	return &Security{
		app: app,
	}
}

func (u *Security) Register(c echo.Context) error  {
	registerUserDto := new(dto.RegisterUserDto)

	if err := c.Bind(registerUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(registerUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := &entity.User{
		Name: registerUserDto.Name,
		Email: registerUserDto.Email,
		Password: registerUserDto.Password,
	}

	err := u.app.RegisterUser(user)

	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "qwe")
}

func (u *Security) Login(c echo.Context) error {
	loginUserDto := new(dto.LoginUserDto)

	if err := c.Bind(loginUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(loginUserDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tokenString, err := u.app.LoginUser(loginUserDto.Email, loginUserDto.Password)

	if err != nil {
		return err
	}

	type Res struct {
		Token string `json:"token"`
	}

	return c.JSON(http.StatusOK, &Res{Token: *tokenString})
}

func(u *Security) SocialRedirect(c echo.Context) error {
	provider := c.Param("provider")

	if provider == "google" {
		googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
		googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
		googleOauthConfig = &oauth2.Config{
			ClientID:     googleClientId,
			ClientSecret: googleClientSecret,
			RedirectURL:  "http://localhost:8000/auth/social/google/success",
			Scopes: []string{
				"profile",
				"email",
			},
			Endpoint: google.Endpoint,
		}

		// TODO: randomize state
		url := googleOauthConfig.AuthCodeURL("state")
		err := c.Redirect(http.StatusFound, url)
		if err != nil {
			return err
		}
	}

	if provider == "facebook" {
		facebookClientId := os.Getenv("FACEBOOK_CLIENT_ID")
		facebookClientSecret := os.Getenv("FACEBOOK_CLIENT_SECRET")
		facebookOauthConfig = &oauth2.Config{
			ClientID: facebookClientId,
			ClientSecret: facebookClientSecret,
			RedirectURL: "http://localhost:8000/auth/social/facebook/success",
			Scopes: []string{
				"public_profile",
				"email",
			},
			Endpoint: facebook.Endpoint,
		}

		url := facebookOauthConfig.AuthCodeURL("state")
		err := c.Redirect(http.StatusFound, url)
		if err != nil {
			return err
		}
	}

	return c.String(200, provider)
}

func(u *Security) SocialLoginSuccess(c echo.Context) error {
	provider := c.Param("provider")
	code := c.QueryParam("code")

	if provider == "google" {
		exchange, err := googleOauthConfig.Exchange(oauth2.NoContext, code)

		if err != nil {
			return fmt.Errorf("code exchange failed: %s", err.Error())
		}

		userInfoUrl := "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + exchange.AccessToken

		userInfo, err := getUserInfo(userInfoUrl)

		if err != nil {
			// TODO: return some predefined error
			return err
		}

		var result map[string]interface{}
		err = json.Unmarshal([]byte(userInfo), &result)

		if err != nil {
			return err
		}

		user, err := u.app.FindUserByEmail(result["email"].(string))

		if err != nil {
			println("errors isnt empty")
			println(err.Error())
		}

		// No user in Db -> create
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			user = &entity.User{
				Email: result["email"].(string),
				Name: result["name"].(string),
				Password: randstr.Hex(16),
			}

			err := u.app.RegisterUser(user)

			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
		}

		token, err := u.app.GenerateToken(*user)

		if err != nil {
			println("token error")
			return err
		}

		type Res struct {
			Token string `json:"token"`
		}

		return c.JSON(200, &Res{Token: *token})
	}

	if provider == "facebook" {
		exchange, err := facebookOauthConfig.Exchange(oauth2.NoContext, code)

		if err != nil {
			return fmt.Errorf("code exchange failed: %s", err.Error())
		}

		userInfoUrl := "https://graph.facebook.com/me?fields=name,first_name,last_name,email&access_token=" + exchange.AccessToken

		userInfo, err := getUserInfo(userInfoUrl)

		if err != nil {
			// TODO: return some predefined error
			return err
		}

		var result map[string]interface{}
		err = json.Unmarshal([]byte(userInfo), &result)

		if err != nil {
			return err
		}

		user, err := u.app.FindUserByEmail(result["email"].(string))

		if err != nil {
			println("errors isnt empty")
			println(err.Error())
		}

		// No user in Db -> create
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			user = &entity.User{
				Email: result["email"].(string),
				Name: result["name"].(string),
				Password: randstr.Hex(16),
			}

			err := u.app.RegisterUser(user)

			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err)
			}
		}

		token, err := u.app.GenerateToken(*user)

		if err != nil {
			println("token error")
			return err
		}

		type Res struct {
			Token string `json:"token"`
		}

		return c.JSON(200, &Res{Token: *token})
	}

	return echo.NewHTTPError(500, "Unsupported Ouath redirect")
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

func(u *Security) UiLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{})
}

func(u *Security) UiRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", map[string]interface{}{})
}

