package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/repositories"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/services"
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
	WithRepository(repository repositories.Repository) Builder
	WithService(service services.Service) Builder
	Now() (Database, error)
}

// Database represents a database instruction
type Database interface {
	Hash() hash.Hash
	IsRepository() bool
	Repository() repositories.Repository
	IsService() bool
	Service() services.Service
}
