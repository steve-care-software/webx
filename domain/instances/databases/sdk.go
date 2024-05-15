package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
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
	WithPath(path []string) Builder
	WithDescription(description string) Builder
	WithHead(head commits.Commit) Builder
	IsActive() Builder
	Now() (Database, error)
}

// Database represents a database
type Database interface {
	Hash() hash.Hash
	Path() []string
	Description() string
	Head() commits.Commit
	IsActive() bool
}

// Repository represents a database repository
type Repository interface {
	List() ([][]string, error)
	Exists(path []string) (*bool, error)
	Retrieve(path []string) (Database, error)
}

// Service represents a database service
type Service interface {
	Save(database Database) error
	Delete(hash hash.Hash) error
}
