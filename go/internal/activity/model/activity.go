package model

import "time"

type CreateIn struct {
	Title string
	Email string
}

type CreateOut struct {
	ID        int64
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
