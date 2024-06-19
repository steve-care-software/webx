package commits

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/queries"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents a commit builder
type Builder interface {
	Create() Builder
	WithQueries(queries queries.Queries) Builder
	WithParent(parent hash.Hash) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Queries() queries.Queries
	HasParent() bool
	Parent() hash.Hash
}

// Repository represents a commit repository
type Repository interface {
	Retrieve(hash hash.Hash) (Commit, error)
}

// Service represents a commit service
type Service interface {
	Save(ins Commit) error
	Delete(hash hash.Hash) error
}
