package application

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"path"
	"path/filepath"
	"runtime"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func Renderer() *TemplateRenderer {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	basePath :=  filepath.Dir(d)

	return &TemplateRenderer{
		templates: template.Must(template.ParseGlob(basePath + "/public/views/*.html")),
	}
}