package db

import (
	"database/sql"
	"fmt"
)

type Connection struct {
	DB *sql.DB
}

func NewConnection(driver, dns string) (*Connection, error) {
	db, err := sql.Open(driver, dns)

	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the db: %w", err)
	}

	return &Connection{DB: db}, nil
}
