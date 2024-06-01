package main

import (
	tmw "gec-applet/templateMiddleware"
	"github.com/labstack/echo/v4"
	"html/template"
	"net/http"
)

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "response", "World")
}

func main() {
	t := &tmw.Template{
		Templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Renderer = t
	e.GET("/hello", Hello)
	e.Logger.Fatal(e.Start(":1323"))
}
