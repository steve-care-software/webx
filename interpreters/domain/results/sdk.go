package results

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the result builder
type Builder interface {
	Create() Builder
	WithValues(values []interface{}) Builder
	WithRemaining(remaining []byte) Builder
	IsValid() Builder
	Now() (Result, error)
}

// Result represents the results
type Result interface {
	IsValid() bool
	HasValues() bool
	Values() []interface{}
	HasRemaining() bool
	Remaining() []byte
}
