package merges

import "github.com/steve-care-software/webx/engine/hashes/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the merge adapter
type Adapter interface {
	ToBytes(ins Merge) ([]byte, error)
	ToInstance(bytes []byte) (Merge, error)
}

// Builder represents a merge builder
type Builder interface {
	Create() Builder
	WithBase(base string) Builder
	WithTop(top string) Builder
	Now() (Merge, error)
}

// Merge represents a merge
type Merge interface {
	Hash() hash.Hash
	Base() string
	Top() string
}
