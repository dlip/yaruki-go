package main

import (
	"io"
	"net/http"

	"github.com/a-h/templ"
	views "github.com/dlip/yaruki-go/pkg/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	component := data.(templ.Component)
	component.Render(c.Request().Context(), w)
	return nil
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{}

	e.Renderer = t
	e.GET("/todos", func(c echo.Context) error {
		component := views.Hello("World")
		return c.Render(http.StatusOK, "", component)
	})

	e.File("/", "public/index.html")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
