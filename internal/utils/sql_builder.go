package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/oliveira533/cubic_ORM.git/internal/db"
	"github.com/oliveira533/cubic_ORM.git/internal/dialects"
)

type SQL_Builder struct {
	dialect dialects.DialectInterface
}

func NewSQL_Builder(dialect dialects.DialectInterface) *SQL_Builder {
	return &SQL_Builder{
		dialect: dialect,
	}
}

func (sql_builder *SQL_Builder) Insert(model any) (string, []any, error) {
	value := reflect.ValueOf(model)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	fields, table := MappingStruct(model)

	var coluns []string
	var placeholders []string
	var args []any

	for idx, field := range fields {
		if sql_builder.hasMeta(field.MataFields, "auto_increment") {
			continue
		}

		coluns = append(coluns, field.ColumnName)
		placeholders = append(placeholders, sql_builder.dialect.Placeholder(len(coluns)))
		// get the value we want insert and convert the generic type in interface to inser in db
		args = append(args, value.Field(idx).Interface())
	}

	var builder strings.Builder
	builder.WriteString("INSERT INTO ")
	builder.WriteString(table)
	builder.WriteString(" (")
	builder.WriteString(strings.Join(coluns, ", "))
	builder.WriteString(") VALUES (")
	builder.WriteString(strings.Join(placeholders, ", "))
	builder.WriteString(")")

	query := builder.String()

	if suffix := sql_builder.dialect.InsertSuffix(); suffix != "" {
		query += " " + suffix
	}

	return query, args, nil

}

func (sql_builder *SQL_Builder) Select(query db.Select) (string, []any, error) {
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

	builder.WriteString("SELECT ")
	builder.WriteString(strings.Join(cols, ", "))
	builder.WriteString(" FROM ")
	builder.WriteString(from)

	var args []any
	if len(query.Where) > 0 {
		builder.WriteString(" WHERE ")
		clauses := make([]string, len(query.Where))

		for idx, clause := range query.Where {
			placeholder := sql_builder.dialect.Placeholder(idx + 1)
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

func (sql_builder *SQL_Builder) Update(query db.Update) (string, []any, error) {
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

	builder.WriteString("UPDATE ")
	builder.WriteString(query.Table)

	var args []any
	if len(query.Where) > 0 {
		clauses := make([]string, len(query.Where))
		builder.WriteString(" WHERE ")
		for idx, clause := range query.Where {
			placeholder := sql_builder.dialect.Placeholder(idx + 1)
			clauses[idx] = fmt.Sprintf("%s %s", clause, placeholder)

			if idx < len(query.Args) {
				args = append(args, query.Args[idx])
			}
		}
		op := "AND"
		if query.Operator != nil {
			op = *query.Operator
		}
		builder.WriteString(strings.Join(clauses, fmt.Sprintf(" %s ", op)))
	}

	return builder.String(), args, nil
}

func (sql_builder *SQL_Builder) SafeDelete(query db.SafeDelete) (string, []any, error) {
	if len(query.Where) == 0 {
		return "", nil, fmt.Errorf("can not delete without where")
	}

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

	builder.WriteString("DELETE FROM ")
	builder.WriteString(from)

	var args []any
	if len(query.Where) > 0 {
		clauses := make([]string, len(query.Where))
		builder.WriteString(" WHERE ")
		for idx, clause := range query.Where {
			placeholder := sql_builder.dialect.Placeholder(idx + 1)
			clauses[idx] = fmt.Sprintf("%s %s", clause, placeholder)

			if idx < len(query.Args) {
				args = append(args, query.Args[idx])
			}
		}
		op := "AND"
		if query.Operator != nil {
			op = *query.Operator
		}
		builder.WriteString(strings.Join(clauses, fmt.Sprintf(" %s ", op)))

	}

	return builder.String(), args, nil
}

func (sql_builder *SQL_Builder) HardDelete(query db.HardDelete) (string, []any, error) {
	from := query.Table

	if from == "" && query.Model != nil {
		_, table := MappingStruct(query.Model)
		from = table
	}

	if from == "" {
		return "", nil, fmt.Errorf("table name or model is required")
	}

	builder := strings.Builder{}
	builder.WriteString("DELETE FROM ")
	builder.WriteString(from)

	var args []any
	if len(query.Where) > 0 {
		clauses := make([]string, len(query.Where))
		builder.WriteString(" WHERE ")
		for idx, clause := range query.Where {
			placeholder := sql_builder.dialect.Placeholder(idx + 1)
			clauses[idx] = fmt.Sprintf("%s %s", clause, placeholder)

			if idx < len(query.Args) {
				args = append(args, query.Args[idx])
			}
		}
		op := "AND"
		if query.Operator != nil {
			op = *query.Operator
		}
		builder.WriteString(strings.Join(clauses, fmt.Sprintf(" %s ", op)))
	}

	return builder.String(), args, nil
}

func (sql_builder *SQL_Builder) hasMeta(meta []db.MetaField, title string) bool {
	for _, m := range meta {
		if m.Title == title {
			return true
		}
	}

	return false
}
