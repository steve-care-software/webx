package tokens

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a token builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithReverse(reverse hash.Hash) Builder
	WithElement(element hash.Hash) Builder
	WithElementIndex(elementIndex uint) Builder
	WithContentIndex(contentIndex uint) Builder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Hash() hash.Hash
	Reverse() hash.Hash
	Element() Element
	HasContent() bool
	Content() *uint
}

// Element represents an element
type Element interface {
	Element() hash.Hash
	Index() uint
}
