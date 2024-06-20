package handler

import (
	"github.com/HergyoBotond/whatToEat/view/recipe"
	"github.com/labstack/echo/v4"
)

type RecipeHandler struct {}

func (h RecipeHandler) HandleRecipeShow(c echo.Context) error {
    return render(c, recipe.Show())
}
