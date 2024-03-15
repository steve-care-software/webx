package logics

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/logics/locations"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents the logic builder
type Builder interface {
	Create() Builder
	WithLayer(layer layers.Layer) Builder
	WithLocation(location locations.Location) Builder
	Now() (Logic, error)
}

// Logic represents logic
type Logic interface {
	Hash() hash.Hash
	Layer() layers.Layer
	Location() locations.Location
}
