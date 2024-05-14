package databases

import "github.com/steve-care-software/datastencil/domain/hash"

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

// Builder represents the database builder
type Builder interface {
	Create() Builder
	WithPath(path string) Builder
	WithDescription(description string) Builder
	WithHead(head string) Builder
	WithActive(isActive string) Builder
	Now() (Database, error)
}

// Database represents a database
type Database interface {
	Hash() hash.Hash
	Path() string
	Description() string
	Head() string
	IsActive() string
}
