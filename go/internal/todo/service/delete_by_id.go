package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"todo/internal/shared/primitive"
)

func (s *Service) DeleteById(ctx context.Context, id int64) (out primitive.BaseResponse) {
	affected, err := s.repo.DeleteById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			out.Status = primitive.ResponseStatusNotFound
			out.SetResponse(http.StatusNotFound, fmt.Sprintf("Todo with ID %d Not Found", id), err)
			return
		}

		out.Status = primitive.ResponseStatusInternalServerError
		out.SetResponse(http.StatusInternalServerError, "internal server error", err)
		return
	}
	if affected == 0 {
		out.Status = primitive.ResponseStatusNotFound
		out.SetResponse(http.StatusNotFound, fmt.Sprintf("Todo with ID %d Not Found", id), err)
		return
	}

	out.Status = primitive.ResponseStatusSuccess
	out.SetResponse(http.StatusOK, "Success")

	return
}
