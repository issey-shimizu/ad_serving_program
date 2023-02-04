package model

import "time"

// Impression ...
type Conversion struct {
	Id            int       `db:"id"`
	Adverrtise_id int       `db:"adverrtise_id"`
	User_code     string    `db:"user_code"`
	Conversion    int       `db:"conversion"`
	Created_at    time.Time `db:"created_at"`
	Updated_at    time.Time `db:"updated_at"`
}
