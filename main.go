package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", helloWorld)

	e.Logger.Fatal(e.Start(":8080"))
}

func helloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, spec.Response{
		Code:    200,
		Message: "Hello World",
	})
}
