package interfaces

import (
	"database/sql"

	"github.com/oliveira533/cubic_ORM.git/internal/db"
)

type Cubic interface {
	Insert(model any) (sql.Result, error)
	Select(query db.Select) (sql.Result, error)
	Update(query db.Update) (sql.Result, error)
}
