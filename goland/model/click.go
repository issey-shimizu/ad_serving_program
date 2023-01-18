package model

import "time"

// click ...
type Click struct {
	id            int       `db:"id"`
	adverrtise_id int       `db:"adverrtise_id"`
	user_code     string    `db:"user_code"`
	click         int       `db:"click"`
	Created_at    time.Time `db:"created_at"`
	Updated_at    time.Time `db:"updated_at"`
}
