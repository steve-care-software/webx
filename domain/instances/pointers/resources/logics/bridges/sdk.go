package bridges

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers"
)

// Builder represents bridges builder
type Builder interface {
	Create() Builder
	WithList(list []Bridge) Builder
	Now() (Bridges, error)
}

// Bridges represents bridges
type Bridges interface {
	Hash() hash.Hash
	List() []Bridge
	Fetch(path []string) (Bridge, error)
}

// BridgeBuilder represents a bridge builder
type BridgeBuilder interface {
	Create() BridgeBuilder
	WithPath(path []string) BridgeBuilder
	WithLayer(layer layers.Layer) BridgeBuilder
	Now() (Bridge, error)
}

// Bridge represents a bridge
type Bridge interface {
	Hash() hash.Hash
	Path() []string
	Layer() layers.Layer
}
