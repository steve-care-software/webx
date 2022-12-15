package insides

import (
	"github.com/steve-care-software/webx/roots/domain/blockchains/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents an inside builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithFn(fn hash.Hash) Builder
	WithFetchers(fetchers []hash.Hash) Builder
	Now() (Inside, error)
}

// Inside represents an inside
type Inside interface {
	Hash() hash.Hash
	Content() Content
}

// Content represents an inside content
type Content interface {
	IsFn() bool
	Fn() hash.Hash
	IsFetchers() bool
	Fetchers() []hash.Hash
}
