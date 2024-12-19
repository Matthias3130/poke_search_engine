package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	etag "github.com/pablor21/echo-etag/v4"
	"main.go/templates"
)

func main() {
	e := echo.New()

	const DB = "pokedex.db"

	conn, err := sql.Open("sqlite3", DB)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := conn.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to SQLite database!")
	defer conn.Close()

	e.Use(etag.Etag())
	e.Static("assets", "./assets")

	e.GET("/", func(ctx echo.Context) error {
		pokemons, err := db.getAll(conn)
		if err != nil {
			log.Fatal(err)
		}
		return Render(ctx, http.StatusOK, templates.Base(templates.SearchBar(pokemons)))
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
