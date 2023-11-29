package dto

import "todo/internal/shared/primitive"

type CreateReq struct {
	Titile string `json:"title"`
	Email  string `json:"email"`
}

type CreateRes struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetByIdRes struct {
	ID        int64            `json:"id"`
	Title     string           `json:"title"`
	Email     string           `json:"email"`
	CreatedAt string           `json:"created_at"`
	UpdatedAt string           `json:"updated_at"`
	DeletedAt primitive.String `json:"deleted_at"`
}
