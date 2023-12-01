package service

import (
	"context"
	"net/http"
	"time"
	"todo/internal/activity/dto"
	"todo/internal/activity/model"
	"todo/internal/shared/primitive"
)

func ValidateCreate(req dto.CreateReq) *primitive.RequestValidationError {
	var allIssues []primitive.RequestValidationIssue

	// validate email
	if len(req.Email) == 0 {
		allIssues = append(allIssues, primitive.RequestValidationIssue{
			Code:    primitive.RequestValidationCodeTooShort,
			Field:   "email",
			Message: "email is required",
		})
	} else {
		if len(req.Email) > 255 {
			allIssues = append(allIssues, primitive.RequestValidationIssue{
				Code:    primitive.RequestValidationCodeTooLong,
				Field:   "email",
				Message: "email too long",
			})
		}

		if !primitive.EmailPattern.MatchString(req.Email) {
			allIssues = append(allIssues, primitive.RequestValidationIssue{
				Code:    primitive.RequestValidationCodeInvalidValue,
				Field:   "email",
				Message: "invalid email",
			})
		}
	}

	// validate title
	if len(req.Title) == 0 {
		allIssues = append(allIssues, primitive.RequestValidationIssue{
			Code:    primitive.RequestValidationCodeTooShort,
			Field:   "title",
			Message: "title cannot be null",
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

	created, err := s.repo.Create(ctx, model.CreateIn{
		Title: req.Title,
		Email: req.Email,
	})
	if err != nil {
		out.Status = primitive.ResponseStatusInternalServerError
		out.SetResponse(http.StatusInternalServerError, "internal server error", err)
		return
	}

	out.Status = primitive.ResponseStatusSuccess
	out.Data = dto.CreateRes{
		ID:        created.ID,
		Title:     created.Title,
		Email:     created.Email,
		CreatedAt: created.CreatedAt.Format(time.RFC3339),
		UpdatedAt: created.UpdatedAt.Format(time.RFC3339),
	}

	out.SetResponse(http.StatusCreated, "Success")

	return
}
