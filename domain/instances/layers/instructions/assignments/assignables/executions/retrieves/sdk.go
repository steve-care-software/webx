package retrieves

import "github.com/steve-care-software/historydb/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a retrieve builder
type Builder interface {
	Create() Builder
	WithContext(context string) Builder
	WithIndex(index string) Builder
	WithReturn(ret string) Builder
	WithLength(length string) Builder
	Now() (Retrieve, error)
}

// Retrieve represents a retrieve instruction
type Retrieve interface {
	Hash() hash.Hash
	Context() string
	Index() string
	Return() string
	HasLength() bool
	Length() string
}
