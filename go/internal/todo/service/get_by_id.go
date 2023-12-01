package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"
	"todo/internal/shared/primitive"
	"todo/internal/todo/dto"
)

func (s *Service) GetById(ctx context.Context, todoId int64) (out primitive.BaseResponse) {
	todo, err := s.repo.GetById(ctx, todoId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			out.Status = primitive.ResponseStatusNotFound
			out.SetResponse(http.StatusNotFound, fmt.Sprintf("Todo with ID %d Not Found", todoId), err)
			return
		} else {
			out.Status = primitive.ResponseStatusInternalServerError
			out.SetResponse(http.StatusInternalServerError, "internal server error", err)
			return
		}
	}

	var data dto.GetByIdRes
	data.ID = todo.ID
	data.ActivityGroupID = todo.ActivityGroupID
	data.Title = todo.Title
	data.IsActive = todo.IsActive
	data.Priority = todo.Priority
	data.CreatedAt = time.Unix(todo.CreatedAt, 0).Format(time.RFC3339)
	data.UpdatedAt = time.Unix(todo.UpdatedAt, 0).Format(time.RFC3339)

	if todo.DeletedAt.Valid {
		data.DeletedAt.Valid = false
		data.DeletedAt.String = time.Unix(todo.DeletedAt.Int64, 0).Format(time.RFC3339)
	}

	out.Status = primitive.ResponseStatusSuccess
	out.Data = data

	out.SetResponse(http.StatusOK, "Success")
	return
}
