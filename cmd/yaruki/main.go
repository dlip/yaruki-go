package main

import (
	"io"
	"net/http"

	"errors"

	"github.com/dlip/yaruki-go/pkg/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	g "github.com/maragudk/gomponents"
)

type Template struct {
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	component, ok := data.(g.Node)
	if !ok {
		return errors.New("Unable to render template")
	}
	component.Render(w)
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
		return c.Render(http.StatusOK, "", views.Navbar())
	})

	e.File("/", "public/index.html")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
