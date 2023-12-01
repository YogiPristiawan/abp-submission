package model

import "todo/internal/shared/primitive"

type CreateIn struct {
	ActivityGroupID int64
	Title           string
}

type CreateOut struct {
	ID              int64
	Title           string
	ActivityGroupID int64
	IsActive        bool
	Priority        primitive.TodoPriority
	CreatedAt       int64
	UpdatedAt       int64
	DeletedAt       primitive.Int64
}

type GetByIdOut struct {
	ID              int64
	Title           string
	ActivityGroupID int64
	IsActive        bool
	Priority        primitive.TodoPriority
	CreatedAt       int64
	UpdatedAt       int64
	DeletedAt       primitive.Int64
}

type FindAllOut struct {
	ID              int64
	Title           string
	ActivityGroupID int64
	IsActive        bool
	Priority        primitive.TodoPriority
	CreatedAt       int64
	UpdatedAt       int64
	DeletedAt       primitive.Int64
}

type UpdateByIdIn struct {
	Title    primitive.String
	IsActive primitive.Bool
}

type UpdateByIdOut struct {
	ID              int64
	Title           string
	ActivityGroupID int64
	IsActive        bool
	Priority        primitive.TodoPriority
	CreatedAt       int64
	UpdatedAt       int64
	DeletedAt       primitive.Int64
}
