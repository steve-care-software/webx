package repositories

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a repository builder
type Builder interface {
	Create() Builder
	WithList(list string) Builder
	WithRetrieve(retrieve string) Builder
	IsSkeleton() Builder
	IsHeight() Builder
	Now() (Repository, error)
}

// Repository represents a repository
type Repository interface {
	Hash() hash.Hash
	IsSkeleton() bool
	IsHeight() bool
	IsList() bool
	List() string
	IsRetrieve() bool
	Retrieve() string
}
