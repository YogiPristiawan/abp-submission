package model

import (
	"time"
	"todo/internal/shared/primitive"
)

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

type GetByIdOut struct {
	ID        int64
	Title     string
	Email     string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt primitive.Int64
}

type FindAllOut struct {
	ID        int64
	Title     string
	Email     string
	CreatedAt int64
	UpdatedAt int64
	DeletedAt primitive.Int64
}
