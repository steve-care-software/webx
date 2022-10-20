package outputs

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an output builder
type Builder interface {
	Create() Builder
	WithValues(values map[string]interface{}) Builder
	WithRemaining(remaining []byte) Builder
	Now() (Output, error)
}

// Output represents a compiled output
type Output interface {
	Values() map[string]interface{}
	HasRemaining() bool
	Remaining() []byte
}
