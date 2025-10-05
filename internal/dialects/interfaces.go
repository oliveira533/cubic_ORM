package dialects

type DialectInterface interface {
	Name() string
	Placeholder(idx int) string
	InsertSuffix() string
}
