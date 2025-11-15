package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/oliveira533/cubic_ORM.git/internal/dialects"
)

func BuildInsertQuery(dialect dialects.DialectInterface, model any) (string, []any, error) {
	value := reflect.ValueOf(model)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	fields, table := MappingStruct(model)

	var coluns []string
	var placeholders []string
	var args []any

	for idx, field := range fields {
		if hasMeta(field.MataFields, "auto_increment") {
			continue
		}

		coluns = append(coluns, field.ColumnName)
		placeholders = append(placeholders, dialect.Placeholder(len(coluns)))
		// get the value we want insert and convert the generic type in interface to inser in db
		args = append(args, value.Field(idx).Interface())
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(coluns, ", "),
		strings.Join(placeholders, ", "),
	)

	if suffix := dialect.InsertSuffix(); suffix != "" {
		query += " " + suffix
	}

	return query, args, nil

}

func BuildSelectQuery(dialect dialects.DialectInterface, model any) {}

func hasMeta(meta []MetaField, title string) bool {
	for _, m := range meta {
		if m.Title == title {
			return true
		}
	}

	return false
}
