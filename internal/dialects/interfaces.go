package dialects

type Dialects interface {
	Name() string
	Placeholder(idx int) string
	InsertSufix() string
}
