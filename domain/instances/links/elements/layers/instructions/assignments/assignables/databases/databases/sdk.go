package databases

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents the database builder
type Builder interface {
	Create() Builder
	WithPath(path string) Builder
	WithDescription(description string) Builder
	WithHead(head string) Builder
	IsActive() Builder
	Now() (Database, error)
}

// Database represents a database
type Database interface {
	Hash() hash.Hash
	Path() string
	Description() string
	Head() string
	IsActive() bool
}
