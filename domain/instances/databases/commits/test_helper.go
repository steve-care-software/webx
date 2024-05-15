package commits

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
)

// NewCommitForTests creates a new commit for tests
func NewCommitForTests(description string, actions actions.Actions) Commit {
	ins, err := NewBuilder().Create().
		WithDescription(description).
		WithActions(actions).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommitWithParentForTests creates a new commit with parent for tests
func NewCommitWithParentForTests(description string, actions actions.Actions, parent hash.Hash) Commit {
	ins, err := NewBuilder().Create().
		WithDescription(description).
		WithActions(actions).
		WithParent(parent).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
