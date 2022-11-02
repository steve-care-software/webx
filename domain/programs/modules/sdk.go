package modules

// ExecuteFn represents the execute func
type ExecuteFn func(input map[string]interface{}) (interface{}, error)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewModuleBuilder creates a new module builder instance
func NewModuleBuilder() ModuleBuilder {
	return createModuleBuilder()
}

// Builder represents a modules builder
type Builder interface {
	Create() Builder
	WithList(list []Module) Builder
	Now() (Modules, error)
}

// Modules represents modules
type Modules interface {
	List() []Module
	Fetch(name []byte) (Module, error)
}

// ModuleBuilder creates a new module builder
type ModuleBuilder interface {
	Create() ModuleBuilder
	WithName(name []byte) ModuleBuilder
	WithFunc(fn ExecuteFn) ModuleBuilder
	Now() (Module, error)
}

// Module represents a module
type Module interface {
	Name() []byte
	Func() ExecuteFn
}
