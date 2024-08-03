package contexts

import (
	"github.com/steve-care-software/webx/engine/databases/hashes/domain/hash"
)

// NewContextForTests creates a new context for tests
func NewContextForTests(identifier uint, executions []hash.Hash) Context {
	ins, err := NewBuilder().Create().
		WithIdentifier(identifier).
		WithExecutions(executions).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewContextWithHeadForTests creates a new context with head for tests
func NewContextWithHeadForTests(identifier uint, head hash.Hash, executions []hash.Hash) Context {
	ins, err := NewBuilder().Create().
		WithIdentifier(identifier).
		WithHead(head).
		WithExecutions(executions).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
