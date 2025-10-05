package dialects

type PostgreSQLDialect struct{}

func (PostgreSQLDialect) Name() string { return "postgres" }

func (PostgreSQLDialect) Placeholder(i int) string {
	return "$1"
}

func (PostgreSQLDialect) InsertSuffix() string {
	return "RETURNING id" // PostgreSQL usa RETURNING
}
