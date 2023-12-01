package presentation

import (
	"context"
	"log"
	"todo/internal/shared/primitive"
	"todo/internal/todo/dto"
)

type IService interface {
	Create(context.Context, dto.CreateReq) primitive.BaseResponse
	GetById(context.Context, int64) primitive.BaseResponse
	FindAll(context.Context, dto.FindAllQuery) primitive.BaseResponseArray
	UpdateById(context.Context, int64, dto.UpdateByIdReq) primitive.BaseResponse
	// DeleteById(context.Context, int64) primitive.BaseResponse
}

type Presentation struct {
	Service IService
}

func New(service IService) *Presentation {
	if service == nil {
		log.Fatal("[x] service required on todo/presentation module")
	}

	return &Presentation{
		Service: service,
	}
}
