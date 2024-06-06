package main

import (
	"gec-applet/templates"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Hello(c echo.Context) error {
	templateResult := templates.Hello("world!!!")
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return templateResult.Render(c.Request().Context(), c.Response().Writer)
}

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/hello", Hello)
	e.Logger.Fatal(e.Start(":1323"))
}
