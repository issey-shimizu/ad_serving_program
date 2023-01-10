package model

import "time"

// Impression ...
type Impression struct {
	id            int       `db:"id"`
	adverrtise_id int       `db:"adverrtise_id"`
	impression    int       `db:"impression"`
	Created_at    time.Time `db:"created_at"`
	Updated_at    time.Time `db:"updated_at"`
}
