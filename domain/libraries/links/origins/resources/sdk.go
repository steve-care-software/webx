package resources

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new origin resource builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the origin resource builder
type Builder interface {
	Create() Builder
	WithLayer(layer hash.Hash) Builder
	WithLayerBytes(layerBytes []byte) Builder
	IsMandatory() Builder
	Now() (Resource, error)
}

// Resource represents an origin resource
type Resource interface {
	Hash() hash.Hash
	Layer() hash.Hash
	IsMandatory() bool
}
