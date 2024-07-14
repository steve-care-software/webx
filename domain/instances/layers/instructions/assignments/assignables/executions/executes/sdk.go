package executes

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the execute builder
type Builder interface {
	Create() Builder
	WithContext(context string) Builder
	WithInput(input inputs.Input) Builder
	WithReturn(ret string) Builder
	WithLayer(layer inputs.Input) Builder
	Now() (Execute, error)
}

// Execute represents an execute
type Execute interface {
	Hash() hash.Hash
	Context() string
	Input() inputs.Input
	Return() string
	HasLayer() bool
	Layer() inputs.Input
}
