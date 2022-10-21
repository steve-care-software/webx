package values

import "github.com/steve-care-software/webx/domain/cryptography/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a value builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithNumber(number byte) Builder
	Now() (Value, error)
}

// Value represents a value
type Value interface {
	Hash() hash.Hash
	Name() string
	Number() byte
}
