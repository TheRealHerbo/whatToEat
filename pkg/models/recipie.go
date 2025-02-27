package models

type Recipie struct {
	Id          int
	Name        string
	Ingredients string
	Directions  string
}

type RecipieForm struct {
	Name        string `form:"name"`
	Ingredients string `form:"ingredients"`
	Directions  string `form:"directions"`
}
