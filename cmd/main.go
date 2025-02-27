package main

import (
	"context"
	"database/sql"
	_ "embed"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"whattoeat/access"
	"whattoeat/pkg/models"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

//go:embed schema.sql
var ddl string

func main() {
	/*
		counter := 2
		recipie1 := models.Recipie{
			Id:          1,
			Name:        "pancake",
			Ingredients: "1. flour 2. milk 3. eggs",
			Directions:  "this is a test",
		}
		recipie2 := models.Recipie{
			Id:          2,
			Name:        "eggs",
			Ingredients: "1. eggs",
			Directions:  "this is a test",
		}

			db := &models.DataBaseMock{
				Recipies: []models.Recipie{recipie1, recipie2},
			}
	*/

	ctx := context.Background()

	db, err := sql.Open("sqlite", "../what-to-eat.db")
	if err != nil {
		log.Fatalln(err)
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatalln(err)
	}

	queries := access.New(db)

	database := &models.DataBase{
		Queries: queries,
	}

	e := echo.New()
	e.Use(middleware.Logger())

	// Serve the ./css folder at the /css URL path
	e.Static("/css", "css")

	e.Renderer = NewTemplate()

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", nil)
	})

	e.GET("/home", func(c echo.Context) error {
		return c.Render(200, "home_page", nil)
	})

	e.GET("/recipies", func(c echo.Context) error {
		list_recipies := database.Get_all(ctx)
		return c.Render(200, "recipies_page", list_recipies)
	})

	e.GET("/add", func(c echo.Context) error {
		return c.Render(200, "add_page", nil)
	})

	e.POST("/new", func(c echo.Context) error {
		form := new(models.RecipieForm)
		if err := c.Bind(form); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		new_recipie := models.Recipie{
			Id:          0,
			Name:        form.Name,
			Ingredients: form.Ingredients,
			Directions:  form.Directions,
		}
		database.Add(ctx, new_recipie)
		return c.Render(200, "recipies_page", database.Get_all(ctx))
	})

	e.GET("/random", func(c echo.Context) error {
		recipie := database.Get_random(ctx)
		c.Response().Header().Set("HX-Push-Url", "/details/"+strconv.Itoa(recipie.Id))
		return c.Render(200, "details_page", recipie)
	})

	e.GET("/details/:id", func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		recipie := database.Get_by_id(ctx, id)
		return c.Render(200, "details_page", recipie)
	})

	e.DELETE("/delete/:id", func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		database.Queries.DeleteRecipie(ctx, id)
		return c.Render(200, "recipies_page", database.Get_all(ctx))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
