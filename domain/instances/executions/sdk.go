package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions/results"
	source_layers "github.com/steve-care-software/datastencil/domain/instances/layers"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewExecutionBuilder creates a new executionbuilder
func NewExecutionBuilder() ExecutionBuilder {
	hashAdapter := hash.NewAdapter()
	return createExecutionBuilder(
		hashAdapter,
	)
}

// Adapter represents the layers adapter
type Adapter interface {
	ToBytes(ins Execution) ([]byte, error)
	ToInstance(bytes []byte) (Execution, error)
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
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithInput(input []byte) ExecutionBuilder
	WithSource(source source_layers.Layer) ExecutionBuilder
	WithResult(result results.Result) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents an executed layer
type Execution interface {
	Hash() hash.Hash
	Input() []byte
	Source() source_layers.Layer
	Result() results.Result
}

// Repository represents an executions repository
type Repository interface {
	RetrieveAll(dbPath []string, hashes []hash.Hash) (Executions, error)
}

// Service represents a service
type Service interface {
	Save(dbPath []string, ins Execution) error
}
