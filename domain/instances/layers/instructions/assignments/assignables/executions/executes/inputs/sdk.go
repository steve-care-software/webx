package inputs

import (
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the input builder
type Builder interface {
	Create() Builder
	WithValue(value string) Builder
	WithPath(path string) Builder
	Now() (Input, error)
}

// Input represents an input
type Input interface {
	Hash() hash.Hash
	IsValue() bool
	Value() string
	IsPath() bool
	Path() string
}
