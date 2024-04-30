package integers

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new integer operator builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the integer operator builder
type Builder interface {
	Create() Builder
	IsSmallerThan() Builder
	IsBiggerThan() Builder
	IsEqual() Builder
	Now() (Integer, error)
}

// Integer represents an integer operator
type Integer interface {
	Hash() hash.Hash
	IsSmallerThan() bool
	IsBiggerThan() bool
	IsEqual() bool
}
