package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	etag "github.com/pablor21/echo-etag/v4"
	"main.go/templates"
)

func main() {
	e := echo.New()
	e.Use(etag.Etag())
	e.Static("assets", "./assets")

	e.GET("/", func(ctx echo.Context) error {
		return Render(ctx, http.StatusOK, templates.Base())
	})

	e.Logger.Fatal(e.Start(":8008"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
