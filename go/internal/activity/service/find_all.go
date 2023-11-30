package service

import (
	"context"
	"net/http"
	"time"
	"todo/internal/activity/dto"
	"todo/internal/shared/primitive"
)

func (s *Service) FindAll(ctx context.Context) (out primitive.BaseResponseArray) {
	activities, err := s.repo.FindAll(ctx)
	if err != nil {
		out.Status = primitive.ResponseStatusInternalServerError
		out.SetResponse(http.StatusInternalServerError, "internal server error", err)
		return
	}

	// map activities
	out.Status = primitive.ResponseStatusSuccess
	for _, activity := range activities {
		out.Data = append(out.Data, dto.FindAll{
			ID:        activity.ID,
			Title:     activity.Title,
			Email:     activity.Email,
			CreatedAt: time.Unix(activity.CreatedAt, 0).Format(time.RFC3339),
			UpdatedAt: time.Unix(activity.UpdatedAt, 0).Format(time.RFC3339),
		})
	}
	out.SetResponse(http.StatusOK, "Success")

	return
}
