package dialects

type MySQLDialect struct{}

func (MySQLDialect) Name() string { return "mysql" }

func (MySQLDialect) Placeholder(i int) string {
	return "?"
}

func (MySQLDialect) InsertSuffix() string {
	return "" // MySQL não usa RETURNING
}
