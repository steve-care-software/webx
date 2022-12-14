package grammars

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hashtrees"
)

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
