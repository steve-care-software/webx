package originals

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an original builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithDescription(description string) Builder
	Now() (Original, error)
}

// Original represents an original namespace
type Original interface {
	Name() string
	Description() string
}
