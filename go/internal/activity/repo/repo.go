package repo

import (
	"database/sql"
	"log"
)

type Repo struct {
	db *sql.DB
}

type Dependency struct {
	DB *sql.DB
}

func New(d *Dependency) *Repo {
	if d.DB == nil {
		log.Fatal("[x] database connection required on activity/repo module")
	}

	return &Repo{
		db: d.DB,
	}
}
