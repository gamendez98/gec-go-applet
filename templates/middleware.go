package templates

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TemplateRenderer struct{}

func (t *TemplateRenderer) Render(w http.ResponseWriter, name string, data interface{}, c echo.Context) error {
	tmpl, ok := data.(templ.Component)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "invalid template data")
	}
	return tmpl.Render(c, w).Render(w)
}
