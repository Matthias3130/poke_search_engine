package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"poke_search_engine/db"
	"poke_search_engine/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	etag "github.com/pablor21/echo-etag/v4"
)

func main() {
	e := echo.New()

	const DB = "db/pokedex.db"
	var err error
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
		return Render(ctx, http.StatusOK, templates.Base(templates.SearchBar()))
	})

	e.GET("/get_all_pokemon", func(ctx echo.Context) error {
		pokemons, err := db.GetAllPokemon(conn)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(http.StatusOK, pokemons)
	})

	e.POST("/find_pokemon", func(ctx echo.Context) error {
		var request struct {
			Find  string   `json:"find"`
			Order bool     `json:"order"`
			Types []string `json:"types"`
			Gens  []int    `json:"gens"`
			Group []bool   `json:"group"`
		}

		if err := ctx.Bind(&request); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		log.Println("Received Types:", request.Types)
		log.Println("Received request:", request)

		pokemons, err := db.FindPokemon(conn, request.Find, request.Order, request.Types, request.Gens, request.Group)
		if err != nil {
			log.Fatal(err)
		}
		return ctx.JSON(http.StatusOK, pokemons)
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
