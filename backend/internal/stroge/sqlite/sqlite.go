package sqlite

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}
