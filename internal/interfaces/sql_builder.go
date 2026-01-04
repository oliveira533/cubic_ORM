package interfaces

import (
	"github.com/oliveira533/cubic_ORM.git/internal/db"
	"github.com/oliveira533/cubic_ORM.git/internal/dialects"
)

type Builder interface {
	Insert(dialect dialects.DialectInterface, model any) (string, []any, error)
	Select(dialect dialects.DialectInterface, query db.Select) (string, []any, error)
	Update(dialect dialects.DialectInterface, query db.Update) (string, []any, error)
	hasMeta(meta []db.MetaField, title string) bool
}
