package grammars

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hashtrees"
)

const minGrammarLength = hash.Size + 8 + 1

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	hashTreeAdapter := hashtrees.NewAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, hashTreeAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a grammar adapter
type Adapter interface {
	ToContent(ins Grammar) ([]byte, error)
	ToGrammar(content []byte) (Grammar, error)
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
