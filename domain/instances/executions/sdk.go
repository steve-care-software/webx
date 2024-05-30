package executions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewExecutionBuilder creates a new execution builder
func NewExecutionBuilder() ExecutionBuilder {
	hashAdapter := hash.NewAdapter()
	return createExecutionBuilder(
		hashAdapter,
	)
}

// Builder represents an executions builder
type Builder interface {
	Create() Builder
	WithList(list []Execution) Builder
	Now() (Executions, error)
}

// Executions represents executions
type Executions interface {
	Hash() hash.Hash
	List() []Execution
	Databases() ([][]string, error)
	Links(basePath []string) ([][]string, error)
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithLogic(logic links.Link) ExecutionBuilder
	WithDatabase(database databases.Database) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents a layer execution
type Execution interface {
	Hash() hash.Hash
	Logic() links.Link
	Database() databases.Database
}
