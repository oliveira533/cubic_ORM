package pkg

import (
	"database/sql"
	"fmt"

	"github.com/oliveira533/cubic_ORM.git/internal/db"
	"github.com/oliveira533/cubic_ORM.git/internal/utils"
)

type Cubic struct {
	connection *db.Connection
	builder    *utils.SQL_Builder
}

func NewCubic(connection *db.Connection) *Cubic {
	return &Cubic{
		connection: connection,
		builder:    utils.NewSQL_Builder(connection.Dialect),
	}
}

func (cubic *Cubic) Insert(model any) (sql.Result, error) {

	command, args, err := cubic.builder.Insert(model)

	if err != nil {
		return nil, fmt.Errorf("cant generate sql query: %e", err)
	}

	results, err := cubic.connection.DB.Exec(command, args...)

	if err != nil {
		return nil, fmt.Errorf("error while executing the insert query \nerror: %e", err)
	}

	return results, nil
}

func (cubic *Cubic) Select(query db.Select) (sql.Result, error) {
	command, args, err := cubic.builder.Select(query)

	if err != nil {
		return nil, fmt.Errorf("cant generate sql query: %e", err)
	}

	results, err := cubic.connection.DB.Exec(command, args...)

	if err != nil {
		return nil, fmt.Errorf("error while executing the insert query \nerror: %e", err)
	}

	return results, nil
}

func (cubic *Cubic) Update(query db.Update) (sql.Result, error) {
	command, args, err := cubic.builder.Update(query)

	if err != nil {
		return nil, fmt.Errorf("cant generate sql query: %e", err)
	}

	results, err := cubic.connection.DB.Exec(command, args...)

	if err != nil {
		return nil, fmt.Errorf("error while executing the insert query \nerror: %e", err)
	}

	return results, nil
}
