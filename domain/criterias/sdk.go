package criterias

import "github.com/steve-care-software/webx/domain/cryptography/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a criteria builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithIndex(index uint) Builder
	WithChild(child Criteria) Builder
	IncludeChannels() Builder
	Now() (Criteria, error)
}

// Criteria represents a criteria
type Criteria interface {
	Hash() hash.Hash
	Name() string
	IncludeChannels() bool
	HasChild() bool
	Child() Criteria
	HasIndex() bool
	Index() *uint
}
