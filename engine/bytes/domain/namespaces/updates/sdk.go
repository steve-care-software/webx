package updates

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an update builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	Now() (Update, error)
}

// Update represents an update namespace
type Update interface {
	Name() string
	Description() string
}
