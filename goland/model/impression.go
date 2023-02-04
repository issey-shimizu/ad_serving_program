package model

import "time"

// Impression ...
type Impression struct {
	Id            int       `db:"id"`
	Adverrtise_id int       `db:"adverrtise_id"`
	Impression    int       `db:"impression"`
	Created_at    time.Time `db:"created_at"`
	Updated_at    time.Time `db:"updated_at"`
}
