package model

import "time"

// Advertise ...
type Advertise struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	Image_url    string    `db:"image_url"`
	Redirect_url string    `db:"redirect_url"`
	Created_at   time.Time `db:"created_at"`
	Updated_at   time.Time `db:"updated_at"`
}
