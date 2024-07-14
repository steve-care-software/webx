package begins

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a begin builder
type Builder interface {
	Create() Builder
	WithPath(path string) Builder
	WithContext(context string) Builder
	Now() (Begin, error)
}

// Begin represents a begin
type Begin interface {
	Path() string
	Context() string
}
