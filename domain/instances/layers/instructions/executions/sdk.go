package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions/merges"
	"github.com/steve-care-software/historydb/domain/databases/commits"
	"github.com/steve-care-software/historydb/domain/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// Builder represents an execution builder
type Builder interface {
	Create() Builder
	WithCommit(commit commits.Commit) Builder
	WithRollback(rollback string) Builder
	WithCancel(cancel string) Builder
	WithMerge(merge merges.Merge) Builder
	Now() (Execution, error)
}

// Execution represents an execution
type Execution interface {
	Hash() hash.Hash
	IsCommit() bool
	Commit() commits.Commit
	IsRollback() bool
	Rollback() string
	IsCancel() bool
	Cancel() string
	IsMerge() bool
	Merge() merges.Merge
}
