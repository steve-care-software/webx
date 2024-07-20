package executions

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/executions/merges"
)

// NewExecutionForTests creates a new execution for tests
func NewExecutionForTests(executable string, content Content) Execution {
	ins, err := NewBuilder().Create().WithExecutable(executable).WithContent(content).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithCommitForTests creates a new execution with commit for tests
func NewContentWithCommitForTests(commit string) Content {
	ins, err := NewContentBuilder().WithCommit(commit).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithRollbackForTests creates a new execution with rollback for tests
func NewContentWithRollbackForTests(rollback string) Content {
	ins, err := NewContentBuilder().WithRollback(rollback).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithCancelForTests creates a new execution with cancel for tests
func NewContentWithCancelForTests(cancel string) Content {
	ins, err := NewContentBuilder().WithCancel(cancel).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithMergeForTests creates a new execution with merge for tests
func NewContentWithMergeForTests(merge merges.Merge) Content {
	ins, err := NewContentBuilder().WithMerge(merge).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
