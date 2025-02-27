package main

import (
	_ "embed"
	"html/template"
	"io"
	"net/http"
	"strconv"

	_ "modernc.org/sqlite"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

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

func main() {
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

	db := &models.DataBase{
		Recipies: []models.Recipie{recipie1, recipie2},
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
		return c.Render(200, "recipies_page", db)
	})

	e.GET("/add", func(c echo.Context) error {
		return c.Render(200, "add_page", nil)
	})

	e.POST("/new", func(c echo.Context) error {
		counter++
		form := new(models.RecipieForm)
		if err := c.Bind(form); err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		new_recipie := models.Recipie{
			Id:          counter,
			Name:        form.Name,
			Ingredients: form.Ingredients,
			Directions:  form.Directions,
		}
		db.Add(new_recipie)
		return c.Render(200, "recipies_page", db)
	})

	e.GET("/random", func(c echo.Context) error {
		recipie := db.Get_random()
		c.Response().Header().Set("HX-Push-Url", "/details/"+strconv.Itoa(recipie.Id))
		return c.Render(200, "details_page", recipie)
	})

	e.GET("/details/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid input")
		}
		recipie := db.Get_by_id(id)
		return c.Render(200, "details_page", recipie)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
