package relationals

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new relational operator builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the relational adapter
type Adapter interface {
	ToBytes(ins Relational) ([]byte, error)
	ToInstance(bytes []byte) (Relational, error)
}

// Builder represents a relational operator builder
type Builder interface {
	Create() Builder
	IsAnd() Builder
	IsOr() Builder
	Now() (Relational, error)
}

// Relational represents a relational operator
type Relational interface {
	Hash() hash.Hash
	IsAnd() bool
	IsOr() bool
}
