package success

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/layers/results/success/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs/kinds"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new success builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the success adapter
type Adapter interface {
	ToBytes(ins Success) ([]byte, error)
	ToInstance(bytes []byte) (Success, error)
}

// Builder represents the success builder
type Builder interface {
	Create() Builder
	WithOutput(output outputs.Output) Builder
	WithKind(kind kinds.Kind) Builder
	Now() (Success, error)
}

// Success represents success result
type Success interface {
	Hash() hash.Hash
	Output() outputs.Output
	Kind() kinds.Kind
}
