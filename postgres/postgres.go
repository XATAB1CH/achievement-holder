package postgres

import "database/sql"

type database struct {
	db *sql.DB
}