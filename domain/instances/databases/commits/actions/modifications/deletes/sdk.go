package deletes

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates the new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithLength(length uint) Builder
	Now() (Delete, error)
}

// Delete represents a delete
type Delete interface {
	Hash() hash.Hash
	Index() uint
	Length() uint
}
