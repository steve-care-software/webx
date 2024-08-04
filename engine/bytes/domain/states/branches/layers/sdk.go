package layers

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/states/branches/layers/pointers"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewLayerBuilder creates a new layer builder
func NewLayerBuilder() LayerBuilder {
	return createLayerBuilder()
}

// Adapter represents a layer adapter
type Adapter interface {
	InstancesToBytes(ins Layers) ([]byte, error)
	BytesToInstances(data []byte) (Layers, []byte, error)
	InstanceToBytes(ins Layer) ([]byte, error)
	BytesToInstance(data []byte) (Layer, []byte, error)
}

// Builder represents a layers builder
type Builder interface {
	Create() Builder
	WithList(list []Layer) Builder
	Now() (Layers, error)
}

// Layers represents layers
type Layers interface {
	List() []Layer
}

// LayerBuilder represents a layer builder
type LayerBuilder interface {
	Create() LayerBuilder
	WithPointers(pointers pointers.Pointers) LayerBuilder
	IsDeleted() LayerBuilder
	Now() (Layer, error)
}

// Layer represents a branch layer
type Layer interface {
	IsDeleted() bool
	HasPointers() bool
	Pointers() pointers.Pointers
}
