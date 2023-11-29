package database

import "errors"

var (
	ErrConnections = errors.New("failed to connect database")
)
