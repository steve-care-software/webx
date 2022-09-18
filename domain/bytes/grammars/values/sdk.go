package values

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a value builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithNumber(number uint8) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Name() string
	Number() uint8
}
