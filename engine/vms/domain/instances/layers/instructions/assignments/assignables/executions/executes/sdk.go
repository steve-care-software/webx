package executes

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the execute adapter
type Adapter interface {
	ToBytes(ins Execute) ([]byte, error)
	ToInstance(bytes []byte) (Execute, error)
}

// Builder represents the execute builder
type Builder interface {
	Create() Builder
	WithContext(context string) Builder
	WithInput(input inputs.Input) Builder
	WithLayer(layer string) Builder
	Now() (Execute, error)
}

// Execute represents an execute
type Execute interface {
	Hash() hash.Hash
	Context() string
	Input() inputs.Input
	HasLayer() bool
	Layer() string
}
