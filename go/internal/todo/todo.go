package todo

import (
	"database/sql"
	"log"
	"todo/internal/todo/presentation"
	"todo/internal/todo/service"
)

type Todo struct {
	Presentation ITodoPresentation
}

type Dependency struct {
	DB *sql.DB
}

func New(d *Dependency) *Todo {
	if d.DB == nil {
		log.Fatal("[x] database connectino required on todo module")
	}

	// init service
	service := service.New(&service.Dependency{
		DB: d.DB,
	})

	// init presentation
	presentation := presentation.New(service)

	return &Todo{
		Presentation: presentation,
	}
}
