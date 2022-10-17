package indexes

import (
	"github.com/steve-care-software/syntax/domain/syntax/criterias"
	"github.com/steve-care-software/syntax/domain/syntax/databases/cryptography/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewIndexBuilder creates a new index builder
func NewIndexBuilder() IndexBuilder {
	hashAdapter := hash.NewAdapter()
	return createIndexBuilder(hashAdapter)
}

// Builder represents an indexes builder
type Builder interface {
	Create() Builder
	WithList(list []Index) Builder
	Now() (Indexes, error)
}

// Indexes represents indexes
type Indexes interface {
	Hash() hash.Hash
	List() []Index
}

// IndexBuilder represents an index builder
type IndexBuilder interface {
	Create() IndexBuilder
	WithName(name string) IndexBuilder
	WithCriteria(criteria criterias.Criteria) IndexBuilder
	Now() (Index, error)
}

// Index represents an index
type Index interface {
	Hash() hash.Hash
	Name() string
	Criteria() criterias.Criteria
}
