package pkg

import (
	"database/sql"
	"fmt"

	"github.com/oliveira533/cubic_ORM.git/internal/db"
	"github.com/oliveira533/cubic_ORM.git/internal/utils"
)

func Insert(model any, conn *db.Connection) (sql.Result, error) {

	command, args, err := utils.BuildInsertQuery(model, conn.Dialect)

	if err != nil {
		return nil, fmt.Errorf("rant generate sql query: %e", err)
	}

	results, err := conn.DB.Exec(command, args...)

	if err != nil {
		return nil, fmt.Errorf("error while executing the insert query \nerror: %e", err)
	}

	return results, nil
}
