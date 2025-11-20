package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/oliveira533/cubic_ORM.git/internal/db"
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

func BuildSelectQuery(dialect dialects.DialectInterface, query db.Select) (string, []any, error) {
	fields, table := MappingStruct(query.Model)

	cols := query.Fields

	if len(cols) == 0 {
		for _, field := range fields {
			cols = append(cols, field.ColumnName)
		}
	}

	from := query.Table

	if from == "" {
		from = table
	}

	builder := strings.Builder{}

	builder.WriteString("SELEC ")
	builder.WriteString(strings.Join(cols, ", "))
	builder.WriteString(" FROM ")
	builder.WriteString(from)

	var args []any
	if len(query.Where) > 0 {
		builder.WriteString(" WHERE ")
		clauses := make([]string, len(query.Where))

		for idx, clause := range query.Where {
			placeholder := dialect.Placeholder(idx + 1)
			clauses[idx] = fmt.Sprintf("%s %s", clause, placeholder)

			if idx < len(query.Args) {
				args = append(args, query.Args[idx])
			}
		}
		builder.WriteString(strings.Join(clauses, " AND "))
	}

	if query.OrderBy != "" {
		builder.WriteString(" ORDER BY ")
		builder.WriteString(query.OrderBy)
	}

	if query.Limit > 0 {
		builder.WriteString(" LIMIT ")
		builder.WriteString(strconv.Itoa(query.Limit))
	}

	return builder.String(), args, nil
}

func hasMeta(meta []MetaField, title string) bool {
	for _, m := range meta {
		if m.Title == title {
			return true
		}
	}

	return false
}
