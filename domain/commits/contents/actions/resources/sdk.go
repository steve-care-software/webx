package resources

import (
	"github.com/steve-care-software/datastencil/domain/commits/contents/actions/resources/instances"
	"github.com/steve-care-software/datastencil/domain/hash"
)

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
