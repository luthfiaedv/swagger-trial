package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/luthfiaedv/swagger-trial/spec"
)

func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, spec.Response{
		Code:    200,
		Message: "Hello World",
	})
}
