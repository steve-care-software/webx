package executions

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions/merges"
	"github.com/steve-care-software/historydb/domain/databases/commits"
)

// NewExecutionWithCommitForTests creates a new execution with commit for tests
func NewExecutionWithCommitForTests(commit commits.Commit) Execution {
	ins, err := NewBuilder().WithCommit(commit).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithRollbackForTests creates a new execution with rollback for tests
func NewExecutionWithRollbackForTests(rollback string) Execution {
	ins, err := NewBuilder().WithRollback(rollback).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithCancelForTests creates a new execution with cancel for tests
func NewExecutionWithCancelForTests(cancel string) Execution {
	ins, err := NewBuilder().WithCancel(cancel).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewExecutionWithMergeForTests creates a new execution with merge for tests
func NewExecutionWithMergeForTests(merge merges.Merge) Execution {
	ins, err := NewBuilder().WithMerge(merge).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
