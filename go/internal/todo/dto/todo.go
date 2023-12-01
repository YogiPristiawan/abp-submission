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
	ID              int64                  `json:"id"`
	Title           string                 `json:"title"`
	ActivityGroupID int64                  `json:"activity_group_id"`
	IsActive        bool                   `json:"is_active"`
	Priority        primitive.TodoPriority `json:"priority"`
	CreatedAt       string                 `json:"created_at"`
	UpdatedAt       string                 `json:"updated_at"`
	DeletedAt       primitive.String       `json:"deleted_at"`
}

type FindAllQuery struct {
	ActivityGroupID int64 `query:"activity_group_id,ommitempty"`
}

type FindAllRes struct {
	ID              int64                  `json:"id"`
	Title           string                 `json:"title"`
	ActivityGroupID int64                  `json:"activity_group_id"`
	IsActive        bool                   `json:"is_active"`
	Priority        primitive.TodoPriority `json:"priority"`
	CreatedAt       string                 `json:"created_at"`
	UpdatedAt       string                 `json:"updated_at"`
	DeletedAt       primitive.Int64        `json:"deleted_at"`
}

type UpdateByIdReq struct {
	Title    primitive.String `json:"title"`
	IsActive primitive.Bool   `json:"is_active"`
}

type UpdateByIdRes struct {
	ID              int64                  `json:"id"`
	Title           string                 `json:"title"`
	ActivityGroupID string                 `json:"activity_group_id"`
	IsActive        string                 `json:"is_active"`
	Priority        primitive.TodoPriority `json:"priority"`
	CreatedAt       string                 `json:"created_at"`
	UpdatedAt       string                 `json:"updated_at"`
	DeletedAt       primitive.String       `json:"deleted_at"`
}
