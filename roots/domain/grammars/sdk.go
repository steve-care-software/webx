package grammars

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hashtrees"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a grammar builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithHistory(history hashtrees.HashTree) Builder
	Now() (Grammar, error)
}

// Grammar represents the grammar database instance
type Grammar interface {
	Hash() hash.Hash
	Name() string
	HasHistory() bool
	History() hashtrees.HashTree
}
