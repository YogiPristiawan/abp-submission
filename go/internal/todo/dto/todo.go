package dto

import "todo/internal/shared/primitive"

type CreateReq struct {
	ActivityGroupID int64  `json:"activity_group_id"`
	Title           string `json:"title"`
}

type CreateRes struct {
	ID              int64                  `json:"id"`
	Title           string                 `json:"title"`
	ActivityGroupID int64                  `json:"activity_group_id"`
	IsActive        bool                   `json:"is_active"`
	Priority        primitive.TodoPriority `json:"priority"`
	CreatedAt       string                 `json:"created_at"`
	UpdatedAt       string                 `json:"updated_at"`
}

type GetByIdRes struct {
	ID              int64
	Title           string
	ActivityGroupID int64
	IsActive        bool
	Priority        primitive.TodoPriority
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       primitive.String
}

type FindAllQuery struct {
	ActivityGroupID int64 `query:"activity_group_id,ommitempty"`
}

type FindAllRes struct {
	ID              int64
	Title           string
	ActivityGroupID int64
	IsActive        bool
	Priority        primitive.TodoPriority
	CreatedAt       string
	UpdatedAt       string
	DeletedAt       primitive.Int64
}
