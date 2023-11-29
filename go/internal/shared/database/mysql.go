package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type MySql struct {
}

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Connect(cfg *MysqlConfig) (*sql.DB, error) {
	if cfg == nil {
		log.Fatal("[x] mysql config is required")
	}

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, ErrConnections
	}

	return db, nil
}
