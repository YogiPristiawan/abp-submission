package activity

import (
	"database/sql"
	"log"
	"todo/internal/activity/presentation"
	"todo/internal/activity/service"
)

type Activity struct {
	Presentation IActivityPresentation
}

type Dependency struct {
	DB *sql.DB
}

func New(d *Dependency) *Activity {
	if d.DB == nil {
		log.Fatal("[x] database connectino required on activity module")
	}

	// init service
	service := service.New(&service.Dependency{
		DB: d.DB,
	})

	// init presentation
	presentation := presentation.New(service)

	return &Activity{
		Presentation: presentation,
	}
}
