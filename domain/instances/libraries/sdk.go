package libraries

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/references"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the library adapter
type Adapter interface {
	ToData(ins Library) ([]byte, error)
	ToInstance(data []byte) (Library, error)
}

// Builder represents the library builder
type Builder interface {
	Create() Builder
	WithLayers(layers layers.Layers) Builder
	WithLinks(links links.Links) Builder
	WithReferences(references references.References) Builder
	Now() (Library, error)
}

// Library represents the library
type Library interface {
	Hash() hash.Hash
	Layers() layers.Layers
	HasLinks() bool
	Links() links.Links
	HasReferences() bool
	References() references.References
}

// Repository represents the library repository
type Repository interface {
	Retrieve(path []string) (Library, error)
}
