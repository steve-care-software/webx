package executions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithLogic(logic links.Link) Builder
	WithDatabase(database databases.Database) Builder
	Now() (Execution, error)
}

// Execution represents a layer execution
type Execution interface {
	Hash() hash.Hash
	Logic() links.Link
	Database() databases.Database
}
