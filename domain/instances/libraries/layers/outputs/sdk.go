package outputs

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/outputs/kinds"
)

// NewBuilder creates a new output builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an output builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithKind(kind kinds.Kind) Builder
	WithExecute(execute []string) Builder
	Now() (Output, error)
}

// Output represents the output
type Output interface {
	Hash() hash.Hash
	Variable() string
	Kind() kinds.Kind
	HasExecute() bool
	Execute() []string
}
