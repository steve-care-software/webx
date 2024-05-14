package commits

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the action adapter
type Adapter interface {
	ToBytes(ins Commit) ([]byte, error)
	ToInstance(bytes []byte) (Commit, error)
}

// Builder represents the commit builder
type Builder interface {
	Create() Builder
	WithDescription(description string) Builder
	WithActions(actions string) Builder
	WithParent(parent string) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Description() string
	Actions() string
	HashParent() bool
	Parent() string
}
