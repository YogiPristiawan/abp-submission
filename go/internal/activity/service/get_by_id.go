package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"
	"todo/internal/activity/dto"
	"todo/internal/shared/primitive"
)

func (s *Service) GetById(ctx context.Context, activityId int64) (out primitive.BaseResponse) {
	activity, err := s.repo.GetById(ctx, activityId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			out.Status = primitive.ResponseStatusNotFound
			out.SetResponse(http.StatusNotFound, fmt.Sprintf("Activity ID with %d Not Found", activityId), err)
			return
		} else {
			out.Status = primitive.ResponseStatusInternalServerError
			out.SetResponse(http.StatusInternalServerError, "internal server error", err)
			return
		}
	}

	var data dto.GetByIdRes
	data.ID = activity.ID
	data.Title = activity.Title
	data.Email = activity.Email
	data.CreatedAt = time.Unix(activity.CreatedAt, 0).Format(time.RFC3339)
	data.UpdatedAt = time.Unix(activity.UpdatedAt, 0).Format(time.RFC3339)
	if activity.DeletedAt.Valid {
		data.DeletedAt.Valid = false
		data.DeletedAt.String = time.Unix(activity.DeletedAt.Int64, 0).Format(time.RFC3339)
	}

	out.Status = primitive.ResponseStatusSuccess
	out.Data = data

	out.SetResponse(http.StatusOK, "Success")
	return
}
