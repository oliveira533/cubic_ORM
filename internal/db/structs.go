package db

import (
	"database/sql"

	"github.com/oliveira533/cubic_ORM.git/internal/dialects"
)

type Connection struct {
	DB      *sql.DB
	Dialect dialects.DialectInterface
}
