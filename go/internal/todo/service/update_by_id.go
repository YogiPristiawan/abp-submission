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
	"todo/internal/todo/model"
)

func ValidateUpdateByIdReq(req dto.UpdateByIdReq) *primitive.RequestValidationError {
	var allIssues []primitive.RequestValidationIssue

	if !req.IsActive.Valid {
		allIssues = append(allIssues, primitive.RequestValidationIssue{
			Code:    primitive.RequestValidationCodeTooShort,
			Field:   "is_active",
			Message: "is_active is required",
		})
	}

	if len(allIssues) > 0 {
		return &primitive.RequestValidationError{
			Issues: allIssues,
		}
	}

	return nil
}

func (s *Service) UpdateById(ctx context.Context, id int64, req dto.UpdateByIdReq) (out primitive.BaseResponse) {
	if err := ValidateUpdateByIdReq(req); err != nil {
		out.Status = primitive.ResponseStatusBadRequest
		out.SetResponse(http.StatusBadRequest, err.ErrorFirst(), err)
		return
	}

	updated, err := s.repo.UpdateById(ctx, id, model.UpdateByIdIn{
		IsActive: req.IsActive.Bool,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			out.Status = primitive.ResponseStatusNotFound
			out.SetResponse(http.StatusNotFound, fmt.Sprintf("Todo with ID %d Not Found", id))
			return
		} else {
			out.Status = primitive.ResponseStatusInternalServerError
			out.SetResponse(http.StatusInternalServerError, "internal server error", err)
			return
		}
	}

	var data dto.UpdateByIdRes
	data.ID = updated.ID
	data.Title = updated.Title
	data.ActivityGroupID = fmt.Sprintf("%d", updated.ActivityGroupID)
	if updated.IsActive {
		data.IsActive = "1"
	} else {
		data.IsActive = "0"
	}
	data.Priority = updated.Priority
	data.CreatedAt = time.Unix(updated.CreatedAt, 0).Format(time.RFC3339)
	data.UpdatedAt = time.Unix(updated.UpdatedAt, 0).Format(time.RFC3339)
	if updated.DeletedAt.Valid {
		data.DeletedAt.Valid = true
		data.DeletedAt.String = time.Unix(updated.DeletedAt.Int64, 0).Format(time.RFC3339)
	}

	out.Status = primitive.ResponseStatusSuccess
	out.Data = data
	out.SetResponse(http.StatusOK, "Success")

	return
}
