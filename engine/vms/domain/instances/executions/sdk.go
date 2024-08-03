package executions

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results"
	source_layers "github.com/steve-care-software/webx/engine/vms/domain/instances/layers"
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
	InstanceToBytes(ins Execution) ([]byte, error)
	BytesToInstance(bytes []byte) (Execution, error)
	InstancesToBytes(ins Executions) ([]byte, error)
	BytesToInstances(bytes []byte) (Executions, error)
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
	Source() source_layers.Layer
	Result() results.Result
	HasInput() bool
	Input() []byte
}
