package main

import (
	"github.com/HergyoBotond/whatToEat/handler"
	"github.com/labstack/echo/v4"
)

func main() {
    app := echo.New()

    recipeHandler := handler.recipeHandler{}

    app.GET("/recipes", recipeHandler.HandleRecipeShow)

    app.Start(":3000")

    app.Logger.Fatal(app.Start(":3000"))
}
