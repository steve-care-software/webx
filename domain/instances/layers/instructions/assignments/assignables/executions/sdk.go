package executions

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithFetch(fetch uint) Builder
	IsLength() Builder
	Now() (Execution, error)
}

// Execution represents an execution instruction
type Execution interface {
	Hash() hash.Hash
	IsLength() bool
	IsFetch() bool
	Fetch() *uint
}
