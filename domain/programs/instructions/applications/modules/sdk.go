package modules

// ExecuteFn represents the execute func
type ExecuteFn func(input map[string]interface{}, currentPath string) (interface{}, error)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder creates a new module builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithFunc(fn ExecuteFn) Builder
	Now() (Module, error)
}

// Module represents a module
type Module interface {
	Name() string
	Func() ExecuteFn
}
