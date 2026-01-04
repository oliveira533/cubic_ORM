package db

import (
	"database/sql"

	"github.com/oliveira533/cubic_ORM.git/internal/dialects"
)

type Connection struct {
	DB      *sql.DB
	Dialect dialects.DialectInterface
}

type Select struct {
	Table   string
	Model   any
	Where   []string
	Args    []string
	OrderBy string
	Limit   int
	Fields  []string
}

type Update struct {
	Table  string
	Model  any
	Fields []string
	Args   []string
	Where  []string
}

type FieldInfo struct {
	Name       string
	ColumnName string
	MataFields []MetaField
	Type       string
}

type MetaField struct {
	Title string
	Value *any
}
