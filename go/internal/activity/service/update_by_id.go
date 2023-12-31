package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"
	"todo/internal/activity/dto"
	"todo/internal/activity/model"
	"todo/internal/shared/primitive"
)

func ValidateUpdateById(req dto.UpdateByIdReq) *primitive.RequestValidationError {
	var allIssues []primitive.RequestValidationIssue

	// validate title
	if len(req.Titile) == 0 {
		allIssues = append(allIssues, primitive.RequestValidationIssue{
			Code:    primitive.RequestValidationCodeTooShort,
			Field:   "title",
			Message: "title cannot be null",
		})
	} else {
		if len(req.Titile) > 255 {
			allIssues = append(allIssues, primitive.RequestValidationIssue{
				Code:    primitive.RequestValidationCodeTooLong,
				Field:   "title",
				Message: "email too long",
			})
		}
	}

	if len(allIssues) > 0 {
		return &primitive.RequestValidationError{
			Issues: allIssues,
		}
	}

	return nil
}

func (s *Service) UpdateById(ctx context.Context, id int64, req dto.UpdateByIdReq) (out primitive.BaseResponse) {
	if err := ValidateUpdateById(req); err != nil {
		out.Status = primitive.ResponseStatusBadRequest
		out.SetResponse(http.StatusBadRequest, err.ErrorFirst(), err)
		return
	}

	// update by id
	updated, err := s.repo.UpdateById(ctx, id, model.UpdateByIdIn{Title: req.Titile})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			out.Status = primitive.ResponseStatusNotFound
			out.SetResponse(http.StatusNotFound, fmt.Sprintf("Activity with ID %d Not Found", id), err)
			return
		}

		out.Status = primitive.ResponseStatusInternalServerError
		out.SetResponse(http.StatusInternalServerError, "internal server error", err)
		return
	}

	out.Status = primitive.ResponseStatusSuccess
	out.Data = dto.UpdateByIdRes{
		ID:        updated.ID,
		Title:     updated.Title,
		Email:     updated.Email,
		CreatedAt: time.Unix(updated.CreatedAt, 0).Format((time.RFC3339)),
		UpdatedAt: time.Unix(updated.UpdatedAt, 0).Format(time.RFC3339),
	}

	out.SetResponse(http.StatusOK, "Success")

	return
}
