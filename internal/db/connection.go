package db

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/oliveira533/cubic_ORM.git/internal/dialects"
)

func NewConnection(driver, dns string) (*Connection, error) {
	db, err := sql.Open(driver, dns)

	dialects := map[string]dialects.DialectInterface{
		"postgres": dialects.PostgreSQLDialect{},
		"mysql":    dialects.PostgreSQLDialect{},
	}

	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the db: %w", err)
	}

	return &Connection{DB: db, Dialect: dialects[detectDriverNameByType(db)]}, nil
}

func detectDriverNameByType(db *sql.DB) string {
	t := fmt.Sprintf("%T", db.Driver())
	switch {
	case strings.Contains(t, "pq.Driver"):
		return "postgres"
	case strings.Contains(t, "pgx"): // covers stdlib/pgx
		return "postgres"
	case strings.Contains(t, "mysql"):
		return "mysql"
	case strings.Contains(t, "sqlite"):
		return "sqlite"
	default:
		return "" // desconhecido
	}
}
