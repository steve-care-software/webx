package links

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a link builder
type Builder interface {
	Create() Builder
	WithInput(input []byte) Builder
	WithSource(source links.Link) Builder
	WithLayers(layers layers.Layers) Builder
	WithNext(next Link) Builder
	Now() (Link, error)
}

// Link represents an executed link
type Link interface {
	Hash() hash.Hash
	Input() []byte
	Source() links.Link
	HasLayers() bool
	Layers() layers.Layers
	HasNext() bool
	Next() Link
}
