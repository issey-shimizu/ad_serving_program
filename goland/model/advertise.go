package model

import "time"

// Article ...
type Article struct {
	ID           int       `db:"id"`
	name         string    `db:name`
	image_url    string    `db:image_url`
	redirect_url string    `db:redirect_url`
	created_at   time.Time `db:created_at`
	updated_at   time.Time `db:update_at`
}
