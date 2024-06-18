package layers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links/layers/results"
	source_layers "github.com/steve-care-software/datastencil/domain/instances/layers"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewLayerBuilder creates a new layer builder
func NewLayerBuilder() LayerBuilder {
	hashAdapter := hash.NewAdapter()
	return createLayerBuilder(
		hashAdapter,
	)
}

// Adapter represents the layers adapter
type Adapter interface {
	ToBytes(ins Layers) ([]byte, error)
	ToInstance(bytes []byte) (Layers, error)
}

// Builder represents the layers builder
type Builder interface {
	Create() Builder
	WithList(list []Layer) Builder
	Now() (Layers, error)
}

// Layers represents executed layers
type Layers interface {
	Hash() hash.Hash
	List() []Layer
}

// LayerBuilder represents a layer builder
type LayerBuilder interface {
	Create() LayerBuilder
	WithInput(input []byte) LayerBuilder
	WithSource(source source_layers.Layer) LayerBuilder
	WithResult(result results.Result) LayerBuilder
	Now() (Layer, error)
}

// Layer represents an executed layer
type Layer interface {
	Hash() hash.Hash
	Input() []byte
	Source() source_layers.Layer
	Result() results.Result
}
