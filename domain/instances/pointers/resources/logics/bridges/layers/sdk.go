package layers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/outputs"
)

// NewBuilder creates a new layer builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the layer adapter
type Adapter interface {
	ToBytes(ins Layer) ([]byte, error)
	ToInstance(bytes []byte) (Layer, error)
}

// Builder represents a layer builder
type Builder interface {
	Create() Builder
	WithInstructions(instructions instructions.Instructions) Builder
	WithOutput(output outputs.Output) Builder
	WithInput(input string) Builder
	Now() (Layer, error)
}

// Layer represents a layer
type Layer interface {
	Hash() hash.Hash
	Instructions() instructions.Instructions
	Output() outputs.Output
	Input() string
}
