package elements

import (
	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

const minElementLength = hash.Size + 8

const (
	valueFlag = iota
	externalFlag
	tokenFlag
	everythingFlag
	recursiveFlag
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	cardinalityAdapter := NewCardinalityAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, cardinalityAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewCardinalityAdapter creates a new cardinality adapter
func NewCardinalityAdapter() CardinalityAdapter {
	builder := NewCardinalityBuilder()
	return createCardinalityAdapter(builder)
}

// NewCardinalityBuilder creates a new cardinality builder
func NewCardinalityBuilder() CardinalityBuilder {
	return createCardinalityBuilder()
}

// Adapter represents an element adapter
type Adapter interface {
	ToContent(ins Element) ([]byte, error)
	ToElement(content []byte) (Element, error)
}

// Builder represents an element builder
type Builder interface {
	Create() Builder
	WithHash(hash hash.Hash) Builder
	WithCardinality(cardinality Cardinality) Builder
	WithValue(value uint8) Builder
	WithExternal(external hash.Hash) Builder
	WithToken(token hash.Hash) Builder
	WithEverything(everything hash.Hash) Builder
	WithRecursive(recursive hash.Hash) Builder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Hash() hash.Hash
	Cardinality() Cardinality
	Content() Content
}

// Content represents an element content
type Content interface {
	IsValue() bool
	Value() *uint8
	IsExternal() bool
	External() *hash.Hash
	IsToken() bool
	Token() *hash.Hash
	IsEverything() bool
	Everything() *hash.Hash
	IsRecursive() bool
	Recursive() *hash.Hash
}

// CardinalityAdapter represents a cardinality adapter
type CardinalityAdapter interface {
	ToContent(ins Cardinality) ([]byte, error)
	ToCardinality(content []byte) (Cardinality, error)
}

// CardinalityBuilder represents a cardinality builder
type CardinalityBuilder interface {
	Create() CardinalityBuilder
	WithMin(min uint) CardinalityBuilder
	WithMax(max uint) CardinalityBuilder
	Now() (Cardinality, error)
}

// Cardinality represents cardinality
type Cardinality interface {
	Min() uint
	HasMax() bool
	Max() *uint
}
