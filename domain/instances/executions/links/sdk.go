package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the link adapter
type Adapter interface {
	ToBytes(ins Link) ([]byte, error)
	ToInstance(bytes []byte) (Link, error)
}

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithInput(input []byte) Builder
	WithSource(source links.Link) Builder
	WithLayers(layers layers.Layers) Builder
	Now() (Link, error)
}

// Link represents an executed link
type Link interface {
	Hash() hash.Hash
	Input() []byte
	Source() links.Link
	HasLayers() bool
	Layers() layers.Layers
}
