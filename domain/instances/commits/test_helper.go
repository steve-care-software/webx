package commits

import (
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
)

// NewCommitForTests creates new commit for tests
func NewCommitForTests(content Content, sig signers.Signature) Commit {
	ins, err := NewBuilder().Create().
		WithContent(content).
		WithSignature(sig).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentForTests creates new content for tests
func NewContentForTests(actions actions.Actions) Content {
	ins, err := NewContentBuilder().Create().
		WithActions(actions).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}

// NewContentWithPreviousForTests creates new content with previous for tests
func NewContentWithPreviousForTests(actions actions.Actions, previous Commit) Content {
	ins, err := NewContentBuilder().Create().
		WithActions(actions).
		WithPrevious(previous).
		Now()

	if err != nil {
		panic(err)
	}

	return ins
}
