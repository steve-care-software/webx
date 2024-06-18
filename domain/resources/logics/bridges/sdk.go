package bridges

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/layers"
)

// NewBuiler creates a new builder
func NewBuiler() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewBridgeBuilder creates a new bridge builder
func NewBridgeBuilder() BridgeBuilder {
	hashAdapter := hash.NewAdapter()
	return createBridgeBuilder(
		hashAdapter,
	)
}

// Adapter represents the bridge adapter
type Adapter interface {
	ToBytes(ins Bridges) ([]byte, error)
	ToInstance(bytes []byte) (Bridges, error)
}

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
