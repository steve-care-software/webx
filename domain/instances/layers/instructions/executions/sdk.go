package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions/merges"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Adapter represents the execution adapter
type Adapter interface {
	ToBytes(ins Execution) ([]byte, error)
	ToInstance(bytes []byte) (Execution, error)
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithCommit(commit string) Builder
	WithRollback(rollback string) Builder
	WithCancel(cancel string) Builder
	WithMerge(merge merges.Merge) Builder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Hash() hash.Hash
	IsCommit() bool
	Commit() string
	IsRollback() bool
	Rollback() string
	IsCancel() bool
	Cancel() string
	IsMerge() bool
	Merge() merges.Merge
}
