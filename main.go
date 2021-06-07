package main

import (
	"github.com/labstack/echo/v4"
	"github.com/luthfiaedv/swagger-trial/handler"
)

func main() {
	e := echo.New()

	e.GET("/", handler.HelloWorld)

	e.Logger.Fatal(e.Start(":8080"))
}
