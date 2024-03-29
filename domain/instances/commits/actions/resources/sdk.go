package resources

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Adapter represents the resource adapter
type Adapter interface {
	ToBytes(ins Resource) ([]byte, error)
	ToInstance(bytes []byte) (Resource, error)
}

// Builder represents a resource builder
type Builder interface {
	Create() Builder
	WithPath(path []string) Builder
	WithInstance(instance instances.Instance) Builder
	Now() (Resource, error)
}

// Resource represents a resource
type Resource interface {
	Hash() hash.Hash
	Path() []string
	Instance() instances.Instance
}
