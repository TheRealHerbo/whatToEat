package models

import (
	"fmt"
	"math/rand"
)

type DataBase struct {
	Recipies []Recipie
}

func (d *DataBase) Add(new_recipie Recipie) {
	d.Recipies = append(d.Recipies, new_recipie)
}

func (d *DataBase) Get_random() Recipie {
	max := len(d.Recipies)
	r := rand.Intn(max-0) - 0

	if r >= max {
		fmt.Println("Random number is out of range")
	}

	return d.Recipies[r]
}

func (d *DataBase) Get_by_id(id int) Recipie {
	for _, recipie := range d.Recipies {
		if recipie.Id == id {
			return recipie
		}
	}
	return Recipie{}
}
