package operators

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new operator builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the operator adapter
type Adapter interface {
	ToBytes(ins Operator) ([]byte, error)
	ToInstance(bytes []byte) (Operator, error)
}

// Builder represents the operator builder
type Builder interface {
	Create() Builder
	IsAnd() Builder
	IsOr() Builder
	IsXor() Builder
	Now() (Operator, error)
}

// Operator represents the operator
type Operator interface {
	Hash() hash.Hash
	IsAnd() bool
	IsOr() bool
	IsXor() bool
}
