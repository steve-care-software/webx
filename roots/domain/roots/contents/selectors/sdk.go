package selectors

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hashtrees"
)

// Adapter represents a selector adapter
type Adapter interface {
	ToSelector(content []byte) (Selector, error)
	ToContent(ins Selector) ([]byte, error)
}

// Builder represents a selector builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithName(name string) Builder
	WithGrammar(grammar hash.Hash) Builder
	WithHistory(history hashtrees.HashTree) Builder
	Now() (Selector, error)
}

// Selector represents a selector database
type Selector interface {
	Hash() hash.Hash
	Name() string
	Grammar() hash.Hash
	HasHistory() bool
	History() hashtrees.HashTree
}
