package service

import (
	"context"
	"database/sql"
	"log"
	"todo/internal/activity/model"
	"todo/internal/activity/repo"
)

type IRepository interface {
	Create(context.Context, model.CreateIn) (model.CreateOut, error)
}

type Service struct {
	repo IRepository
}

type Dependency struct {
	DB *sql.DB
}

func New(d *Dependency) *Service {
	if d.DB == nil {
		log.Fatal("[x] database connection required on activity/service module")
	}

	repo := repo.New(&repo.Dependency{
		DB: d.DB,
	})

	return &Service{
		repo: repo,
	}
}