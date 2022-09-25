package outputs

// NewBuilder creates a new buiklder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewVariableBuilder creates a new variable builder
func NewVariableBuilder() VariableBuilder {
	return createVariableBuilder()
}

// Builder represents an output builder
type Builder interface {
	Create() Builder
	WithList(list []Variable) Builder
	Now() (Output, error)
}

// Output represents an output
type Output interface {
	List() []Variable
	Find(name string) (Variable, error)
}

// VariableBuilder represents a variable builder
type VariableBuilder interface {
	Create() VariableBuilder
	WithName(name string) VariableBuilder
	WithValue(value interface{}) VariableBuilder
	Now() (Variable, error)
}

// Variable represents a variable
type Variable interface {
	Name() string
	HasValue() bool
	Value() interface{}
}
