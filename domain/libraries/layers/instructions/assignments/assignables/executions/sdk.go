package executions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new execution builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithInput(input string) Builder
	WithLayer(layer string) Builder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Hash() hash.Hash
	Input() string
	HasLayer() bool
	Layer() string
}
