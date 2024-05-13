package deletes

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

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
