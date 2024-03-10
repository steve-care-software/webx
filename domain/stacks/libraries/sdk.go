package libraries

import "github.com/steve-care-software/datastencil/domain/commits/actions/resources/instances"

// Builder represents a library builder
type Builder interface {
	Create() Builder
	WithInstance(ins instances.Instance) Builder
	Now() (Library, error)
}

// Library represents a library assignable
type Library interface {
	IsInstance() bool
	Instance() instances.Instance
}
