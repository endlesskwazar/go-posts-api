package errors

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	config2 "go-cource-api/application/config"
	"net/http"
)

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	config := c.Get("config").(config2.Config)

	var serverErrorMessage string

	if config.AppConfig.Env == "prod" {
		serverErrorMessage = "Oops! Something went wrong, Please, try again later or contact support"
	} else {
		serverErrorMessage = err.Error()
	}

	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, serverErrorMessage)
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s is not valid email",
					err.Field())
			}

			break
		}
	}

	c.Logger().Error(report)
	err = c.JSON(report.Code, report)
}
