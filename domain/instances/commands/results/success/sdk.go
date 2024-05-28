package success

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results/success/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs/kinds"
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
