package heads

import "github.com/steve-care-software/historydb/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the head adapter
type Adapter interface {
	ToBytes(ins Head) ([]byte, error)
	ToInstance(bytes []byte) (Head, error)
}

// Builder represents the head builder
type Builder interface {
	Create() Builder
	WithContext(context string) Builder
	WithReturn(ret string) Builder
	Now() (Head, error)
}

// Head represents a head
type Head interface {
	Hash() hash.Hash
	Context() string
	Return() string
}
