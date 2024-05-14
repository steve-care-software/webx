package deletes

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a delete builder
type Builder interface {
	Create() Builder
	WithIndex(index string) Builder
	WithLength(length string) Builder
	Now() (Delete, error)
}

// Delete represents a delete
type Delete interface {
	Hash() hash.Hash
	Index() string
	Length() string
}