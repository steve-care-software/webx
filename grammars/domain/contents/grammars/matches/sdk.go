package matches

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

const minMatchSize = hash.Size * 3

// NewAdapter creates a new adapter for tests
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, builder)
}

// NewBuilder creates a new builder for tests
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a match adapter
type Adapter interface {
	ToContent(ins Match) ([]byte, error)
	ToMatch(content []byte) (Match, error)
}

// Builder represents a match builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithToken(token hash.Hash) Builder
	WithSuites(suites []hash.Hash) Builder
	Now() (Match, error)
}

// Match represents a match between a token and a suites
type Match interface {
	Hash() hash.Hash
	Token() hash.Hash
	Suites() []hash.Hash
}
