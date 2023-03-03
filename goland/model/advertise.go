package model

import "time"

// Advertise ...
/* 変更1
type Advertise struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	Image_url    string    `db:"image_url"`
	Redirect_url string    `db:"redirect_url"`
	Created_at   time.Time `db:"created_at"`
	Updated_at   time.Time `db:"updated_at"`
}
*/

type Advertise struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Image_url    string    `json:"image_url"`
	Redirect_url string    `json:"redirect_url"`
	Created_at   time.Time `json:"created_at"`
}
