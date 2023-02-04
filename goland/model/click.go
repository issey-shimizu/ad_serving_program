package model

import "time"

// click ...
type Click struct {
	Id            int       `db:"id"`
	Adverrtise_id int       `db:"adverrtise_id"`
	User_code     string    `db:"user_code"`
	Click         int       `db:"click"`
	Created_at    time.Time `db:"created_at"`
	Updated_at    time.Time `db:"updated_at"`
}
