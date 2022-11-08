package models

import "time"

type ShortUrl struct {
	Id        int       `db:"ID"`
	FullUrl   string    `db:"full_url"`
	Url       string    `db:"url"`
	Click     int       `db:"click"`
	CreatedAt time.Time `json:"created_at"`
}
