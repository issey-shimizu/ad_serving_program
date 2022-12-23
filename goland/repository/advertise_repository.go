package repository

import (
	"log"
	"src/model"
)

// ArticleList ...
func Adverdisplay() ([]*model.Advertise, error) {
	query := `SELECT * FROM advertise;`

	var advertise []*model.Advertise
	if err := db.Select(&advertise, query); err != nil {
		return nil, err
	}
	log.Println(advertise[0].ID)

	type Profile struct {
		Name string
		Age  int
	}

	p := []*Profile{
		{"Tanaka", 31},
		{"Suzuki", 46},
	}
	log.Println(p[0].Name)
	log.Println(p[0].Name)

	return advertise, nil
}
