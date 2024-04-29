package reverts

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the revert adapter
type Adapter interface {
	ToBytes(ins Revert) ([]byte, error)
	ToInstance(bytes []byte) (Revert, error)
}

// Builder represents a revert builder
type Builder interface {
	Create() Builder
	WithIndex(index string) Builder
	Now() (Revert, error)
}

// Revert represents a revert
type Revert interface {
	Hash() hash.Hash
	HasIndex() bool
	Index() string
}
