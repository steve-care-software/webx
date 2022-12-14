package selectors

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a selector builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithGrammar(grammar hash.Hash) Builder
	WithToken(token hash.Hash) Builder
	WithInside(inside hash.Hash) Builder
	WithFunc(fn hash.Hash) Builder
	Now() (Selector, error)
}

// Selector represents a selector
type Selector interface {
	Hash() hash.Hash
	Grammar() hash.Hash
	Token() hash.Hash
	Inside() hash.Hash
	Func() hash.Hash
}
