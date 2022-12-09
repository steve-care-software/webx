package assignments

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an assignment builder
type Builder interface {
	Create() Builder
	WithIndex(index uint) Builder
	WithValue(value hash.Hash) Builder
	Now() (Assignment, error)
}

// Assignment represents an assignment
type Assignment interface {
	Index() uint
	Value() hash.Hash
}
