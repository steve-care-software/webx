package contents

import (
	"github.com/steve-care-software/webx/domain/databases/references"
)

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	Now() (Application, error)
}

// Application represents a file application
type Application interface {
	Reference() (references.Reference, error)
	Retrieve(pointer references.Pointer) ([]byte, error)
	List(pointers []references.Pointer) ([][]byte, error)
}
