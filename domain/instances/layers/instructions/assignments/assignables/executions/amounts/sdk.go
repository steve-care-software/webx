package amounts

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the amount adapter
type Adapter interface {
	ToBytes(ins Amount) ([]byte, error)
	ToInstance(bytes []byte) (Amount, error)
}

// Builder represents the amount instruction
type Builder interface {
	Create() Builder
	WithContext(context string) Builder
	WithReturn(ret string) Builder
	Now() (Amount, error)
}

// Amount represents an amount
type Amount interface {
	Hash() hash.Hash
	Context() string
	Return() string
}
