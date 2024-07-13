package layers

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/layers/results"
	source_layers "github.com/steve-care-software/datastencil/domain/instances/layers"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the layers adapter
type Adapter interface {
	ToBytes(ins Layer) ([]byte, error)
	ToInstance(bytes []byte) (Layer, error)
}

// Builder represents a layer builder
type Builder interface {
	Create() Builder
	WithInput(input []byte) Builder
	WithSource(source source_layers.Layer) Builder
	WithResult(result results.Result) Builder
	Now() (Layer, error)
}

// Layer represents an executed layer
type Layer interface {
	Hash() hash.Hash
	Input() []byte
	Source() source_layers.Layer
	Result() results.Result
}

// Service represents a service
type Service interface {
	Save(ins Layer) error
}
