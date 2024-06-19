package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the database adapter
type Adapter interface {
	ToBytes(ins Database) ([]byte, error)
	ToInstance(bytes []byte) (Database, error)
}

// Builder represents a database builder
type Builder interface {
	Create() Builder
	WithCommit(commit commits.Commit) Builder
	WithHead(head heads.Head) Builder
	Now() (Database, error)
}

// Database represents a database
type Database interface {
	Hash() hash.Hash
	Commit() commits.Commit
	Head() heads.Head
}

// Repository represents a database repository
type Repository interface {
	Retrieve(path []string) (Database, error)
}

// Service represents a database service
type Service interface {
	Save(database Database) error
	Delete(hash hash.Hash) error
}
