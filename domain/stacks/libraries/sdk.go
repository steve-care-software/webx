package libraries

import "github.com/steve-care-software/datastencil/domain/orms"

// Builder represents a library builder
type Builder interface {
	Create() Builder
	WithInstance(ins orms.Instance) Builder
	Now() (Library, error)
}

// Library represents a library assignable
type Library interface {
	IsInstance() bool
	Instance() orms.Instance
}
