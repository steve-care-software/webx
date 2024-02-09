package libraries

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/links"
	"github.com/steve-care-software/identity/domain/hash"
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
	Now() (Library, error)
}

// Library represents the library
type Library interface {
	Hash() hash.Hash
	Layers() layers.Layers
	HasLinks() bool
	Links() links.Links
}

// Repository represents the library repository
type Repository interface {
	Retrieve(path []string) (Library, error)
}
