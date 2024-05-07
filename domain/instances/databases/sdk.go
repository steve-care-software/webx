package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
)

// Adapter represents the database adapter
type Adapter interface {
	ToBytes(ins Database) ([]byte, error)
	ToInstance(data []byte) (Database, error)
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
	List() [][]string
	Retrieve(path []string) (Database, error)
}

// Service represents a database service
type Service interface {
	Insert(database Database) error
	Update(origin hash.Hash, updated Database) error
	Delete(hash hash.Hash) error
	Purge(hash hash.Hash) error
}
