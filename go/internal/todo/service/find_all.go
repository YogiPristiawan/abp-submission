package service

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"todo/internal/shared/primitive"
	"todo/internal/todo/dto"
)

func (s *Service) FindAll(ctx context.Context, query dto.FindAllQuery) (out primitive.BaseResponseArray) {
	var queryFilter []int64
	if query.ActivityGroupID != 0 {
		queryFilter = append(queryFilter, query.ActivityGroupID)
	}

	todos, err := s.repo.FindAll(ctx, queryFilter...)
	if err != nil {
		out.Status = primitive.ResponseStatusInternalServerError
		out.SetResponse(http.StatusInternalServerError, "internal server error", err)
		return
	}

	// map activities
	out.Status = primitive.ResponseStatusSuccess
	if len(out.Data) == 0 {
		fmt.Println("masuk")
		out.Data = []any{}
	} else {
		for _, todo := range todos {
			out.Data = append(out.Data, dto.FindAllRes{
				ID:              todo.ID,
				Title:           todo.Title,
				ActivityGroupID: todo.ActivityGroupID,
				IsActive:        todo.IsActive,
				Priority:        todo.Priority,
				CreatedAt:       time.Unix(todo.CreatedAt, 0).Format(time.RFC3339),
				UpdatedAt:       time.Unix(todo.UpdatedAt, 0).Format(time.RFC3339),
			})
		}
	}
	out.SetResponse(http.StatusOK, "Success")

	return
}
