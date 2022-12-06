package suites

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

const minSuiteSize = hash.Size + 2

// NewAdapter creates a new suite adapter
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, builder)
}

// NewBuilder creates a new suite builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a suite adapter
type Adapter interface {
	ToContent(ins Suite) ([]byte, error)
	ToSuite(content []byte) (Suite, error)
}

// Builder represents a suite builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithContent(content []byte) Builder
	IsValid() Builder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Hash() hash.Hash
	IsValid() bool
	Content() []byte
}
