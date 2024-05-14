package commits

import (
	"time"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
)

// NewCommitForTests creates a new commit for tests
func NewCommitForTests(description string, actions actions.Actions, createdOn time.Time) Commit {
	ins, err := NewBuilder().Create().
		WithDescription(description).
		WithActions(actions).
		CreatedOn(createdOn).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewCommitWithParentForTests creates a new commit with parent for tests
func NewCommitWithParentForTests(description string, actions actions.Actions, createdOn time.Time, parent hash.Hash) Commit {
	ins, err := NewBuilder().Create().
		WithDescription(description).
		WithActions(actions).
		CreatedOn(createdOn).
		WithParent(parent).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
