package elements

import (
	"github.com/steve-care-software/webx/domain/cryptography/hash"
	"github.com/steve-care-software/webx/domain/databases/grammars/cardinalities"
)

const (
	valueFlag = iota
	external
	token
	everything
	recursive
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
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
	WithCardinality(cardinality cardinalities.Cardinality) Builder
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
	Cardinality() cardinalities.Cardinality
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
