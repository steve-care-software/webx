package executions

import "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions/merges"

// Execution represents an execution
type Execution interface {
	IsCommit() bool
	Commit() string
	IsRollback() bool
	Rollback() string
	IsCancel() bool
	Cancel() string
	IsMerge() bool
	Merge() merges.Merge
}
