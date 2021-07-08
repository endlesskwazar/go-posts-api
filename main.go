package main

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"go-cource-api/infrustructure/persistence"
	"go-cource-api/infrustructure/validation"
	"go-cource-api/interfaces/handlers"
	"go-cource-api/interfaces/middlewares"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"io/ioutil"
	http "net/http"
	"os"
)

var (
	googleOauthConfig *oauth2.Config
)

func main() {
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	services, err := persistence.NewRepositories(user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	services.Automigrate()

	posts := handlers.NewPosts(services.Post)
	users := handlers.NewUsers(services.User)

	e := echo.New()
	e.Validator = &validation.CustomValidator{
		Validator: validator.New(),
	}

	e.Renderer = Renderer()

	e.POST("/register", users.Register)
	e.POST("/login", users.Login)

	e.POST("/api/posts", posts.Save)
	e.GET("/api/posts", posts.List)

	e.GET("/api/users", users.List)
	e.POST("/api/users", users.Save)

	r := e.Group("/api")
	r.Use(middlewares.AuthMiddleware())
	r.GET("/restricted", func(context echo.Context) error {
		return context.String(200, "qweqwe")
	})

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{})
	})

	e.GET("/register", func(c echo.Context) error {
		return c.Render(http.StatusOK, "register.html", map[string]interface{}{})
	})

	e.GET("/security/social", func(c echo.Context) error {
		provider := c.QueryParam("provider")

		if provider == "google" {
			googleOauthConfig = &oauth2.Config{
				ClientID:     googleClientId,
				ClientSecret: googleClientSecret,
				RedirectURL:  "http://localhost:8000/security/social/success",
				Scopes: []string{
					"https://www.googleapis.com/security/userinfo.profile",
					"https://www.googleapis.com/security/userinfo.email",
				},
				Endpoint: google.Endpoint,
			}

			// TODO: randomize state
			url := googleOauthConfig.AuthCodeURL("state")
			err := c.Redirect(http.StatusFound, url)
			if err != nil {
				return err
			}
			return nil
		}

		return c.String(200, provider)
	})

	e.GET("/security/social/success", func(c echo.Context) error {
		code := c.QueryParam("code")

		userInfo, err := getUserInfo(code)

		if err != nil {
			// TODO: return some predefined error
			return err
		}

		println(userInfo)

		return c.String(200, string(userInfo))
	})

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func getUserInfo(code string) ([]byte, error) {
	exchange, err := googleOauthConfig.Exchange(oauth2.NoContext, code)

	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + exchange.AccessToken)

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
