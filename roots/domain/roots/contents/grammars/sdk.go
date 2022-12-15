package grammars

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hashtrees"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a grammar adapter
type Adapter interface {
	ToGrammar(content []byte) (Grammar, error)
	ToContent(ins Grammar) ([]byte, error)
}

// Builder represents a grammar builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
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
