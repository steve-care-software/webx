package fetchers

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a fetcher builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithRecursive(recursive hash.Hash) Builder
	WithSelector(selector hash.Hash) Builder
	Now() (Fetcher, error)
}

// Fetcher represents a fetcher
type Fetcher interface {
	Hash() hash.Hash
	Content() Content
}

// Content represents a fetcher's content
type Content interface {
	IsRecursive() bool
	Recursive() *hash.Hash
	IsSelector() bool
	Selector() *hash.Hash
}
