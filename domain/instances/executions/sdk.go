package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/layers"
	"github.com/steve-care-software/historydb/domain/databases"
	"github.com/steve-care-software/historydb/domain/hash"
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

// Adapter represents the executions adapter
type Adapter interface {
	ToBytes(ins Executions) ([]byte, error)
	ToInstance(bytes []byte) (Executions, error)
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
	WithLayer(layer layers.Layer) ExecutionBuilder
	WithDatabase(database databases.Database) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents a layer execution
type Execution interface {
	Hash() hash.Hash
	Layer() layers.Layer
	Database() databases.Database
}
