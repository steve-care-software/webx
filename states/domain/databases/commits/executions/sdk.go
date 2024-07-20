package executions

import (
	"github.com/steve-care-software/datastencil/states/domain/databases/commits/executions/chunks"
	"github.com/steve-care-software/datastencil/states/domain/hash"
)

// NewBuilder creates a new builder
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
}

// ExecutionBuilder represents an execution builder
type ExecutionBuilder interface {
	Create() ExecutionBuilder
	WithBytes(bytes []byte) ExecutionBuilder
	WithChunk(chunk chunks.Chunk) ExecutionBuilder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Hash() hash.Hash
	IsBytes() bool
	Bytes() []byte
	IsChunk() bool
	Chunk() chunks.Chunk
}
