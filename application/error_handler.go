package application

import (
	"github.com/labstack/echo/v4"
	"go-cource-api/application/config"
	"net/http"
)

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	appConfig := c.Get("config").(*config.Config)

	var serverErrorMessage string

	if appConfig.AppConfig.Env == "prod" {
		serverErrorMessage = "Oops! Something went wrong, Please, try again later or contact support"
	} else {
		serverErrorMessage = err.Error()
	}

	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, serverErrorMessage)
	}

	c.Logger().Error(report)
	_ = c.JSON(report.Code, report)
}
