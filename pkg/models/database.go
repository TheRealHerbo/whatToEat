package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"whattoeat/access"
)

type DataBase struct {
	Queries *access.Queries
}

func (d *DataBase) Get_all(ctx context.Context) []Recipie {
	data := []Recipie{}
	recipies, err := d.Queries.ListRecipies(ctx)
	if err != nil {
		log.Fatalln(err)
		return []Recipie{}
	}
	log.Println(recipies)
	for _, r := range recipies {
		data = append(data, Recipie{
			Id:          int(r.ID),
			Name:        r.Name,
			Ingredients: r.Ingredients.String,
			Directions:  r.Directions.String,
		})
	}
	return data
}

func (d *DataBase) Add(ctx context.Context, new_recipie Recipie) error {
	inserted_author, err := d.Queries.CreateRecipie(ctx, access.CreateRecipieParams{
		Name:        new_recipie.Name,
		Ingredients: sql.NullString{String: new_recipie.Ingredients, Valid: true},
		Directions:  sql.NullString{String: new_recipie.Directions, Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(inserted_author)
	return nil
}

func (d *DataBase) Get_random(ctx context.Context) Recipie {
	max, err := d.Queries.CountRecipies(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	r := rand.Int63n(max-0) - 0

	if r >= max {
		fmt.Println("Random number is out of range")
	}
	fetchedRecipie, err := d.Queries.GetRecipie(ctx, r)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fetchedRecipie)

	recipie := Recipie{
		Id:          int(fetchedRecipie.ID),
		Name:        fetchedRecipie.Name,
		Ingredients: fetchedRecipie.Ingredients.String,
		Directions:  fetchedRecipie.Directions.String,
	}

	return recipie
}

func (d *DataBase) Get_by_id(ctx context.Context, id int64) Recipie {
	fetchedRecipie, err := d.Queries.GetRecipie(ctx, id)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fetchedRecipie)

	recipie := Recipie{
		Id:          int(fetchedRecipie.ID),
		Name:        fetchedRecipie.Name,
		Ingredients: fetchedRecipie.Ingredients.String,
		Directions:  fetchedRecipie.Directions.String,
	}

	return recipie
}
