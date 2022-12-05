package suites

import (
	"github.com/steve-care-software/webx/blockchains/domain/cryptography/hash"
)

const minSuiteSize = hash.Size + 2

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	builder := NewBuilder()
	suiteAdapter := NewSuiteAdapter()
	return createAdapter(builder, suiteAdapter)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewSuiteAdapter creates a new suite adapter
func NewSuiteAdapter() SuiteAdapter {
	hashAdapter := hash.NewAdapter()
	builder := NewSuiteBuilder()
	return createSuiteAdapter(hashAdapter, builder)
}

// NewSuiteBuilder creates a new suite builder
func NewSuiteBuilder() SuiteBuilder {
	return createSuiteBuilder()
}

// Adapter represents a suites adapter
type Adapter interface {
	ToContent(ins Suites) ([]byte, error)
	ToSuites(content []byte) (Suites, error)
}

// Builder represents a suites builder
type Builder interface {
	Create() Builder
	WithList(list []Suite) Builder
	Now() (Suites, error)
}

// Suites represents suites
type Suites interface {
	List() []Suite
}

// SuiteAdapter represents a suite adapter
type SuiteAdapter interface {
	ToContent(ins Suite) ([]byte, error)
	ToSuite(content []byte) (Suite, error)
}

// SuiteBuilder represents a suite builder
type SuiteBuilder interface {
	Create() SuiteBuilder
	WithHash(hash hash.Hash) SuiteBuilder
	WithContent(content []byte) SuiteBuilder
	IsValid() SuiteBuilder
	Now() (Suite, error)
}

// Suite represents a suite
type Suite interface {
	Hash() hash.Hash
	IsValid() bool
	Content() []byte
}
