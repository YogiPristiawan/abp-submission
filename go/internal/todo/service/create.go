package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"
	"todo/internal/shared/primitive"
	"todo/internal/todo/dto"
	"todo/internal/todo/model"
)

func ValidateCreate(req dto.CreateReq) *primitive.RequestValidationError {
	var allIssues []primitive.RequestValidationIssue

	// validate email
	if req.ActivityGroupID == 0 {
		allIssues = append(allIssues, primitive.RequestValidationIssue{
			Code:    primitive.RequestValidationCodeTooShort,
			Field:   "activity_group_id",
			Message: "activity_gorup_id is required",
		})
	}

	// validate title
	if len(req.Title) == 0 {
		allIssues = append(allIssues, primitive.RequestValidationIssue{
			Code:    primitive.RequestValidationCodeTooShort,
			Field:   "title",
			Message: "title is required",
		})
	} else {
		if len(req.Title) > 255 {
			allIssues = append(allIssues, primitive.RequestValidationIssue{
				Code:    primitive.RequestValidationCodeTooLong,
				Field:   "title",
				Message: "email too lonng",
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

func (s *Service) Create(ctx context.Context, req dto.CreateReq) (out primitive.BaseResponse) {
	if err := ValidateCreate(req); err != nil {
		out.Status = primitive.ResponseStatusBadRequest
		out.SetResponse(http.StatusBadRequest, err.ErrorFirst(), err)
		return
	}

	// create todo
	created, err := s.repo.Create(ctx, model.CreateIn{
		ActivityGroupID: req.ActivityGroupID,
		Title:           req.Title,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			out.Status = primitive.ResponseStatusNotFound
			out.SetResponse(http.StatusNotFound, "Not Found", err)
			return
		}

		out.Status = primitive.ResponseStatusInternalServerError
		out.SetResponse(http.StatusInternalServerError, "internal server error", err)
		return
	}

	out.Status = primitive.ResponseStatusSuccess
	out.Data = dto.CreateRes{
		ID:              created.ID,
		Title:           created.Title,
		ActivityGroupID: created.ActivityGroupID,
		IsActive:        created.IsActive,
		Priority:        created.Priority,
		CreatedAt:       time.Unix(created.CreatedAt, 0).Format(time.RFC3339),
		UpdatedAt:       time.Unix(created.UpdatedAt, 0).Format(time.RFC3339),
	}
	out.SetResponse(http.StatusCreated, "Success")

	return
}
