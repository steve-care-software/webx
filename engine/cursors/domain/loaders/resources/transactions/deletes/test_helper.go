package deletes

import "github.com/steve-care-software/webx/engine/cursors/domain/loaders/identities/keys/signers"

// NewDeleteForTests creates a new delete for tests
func NewDeleteForTests(name string, vote signers.Vote) Delete {
	ins, err := NewBuilder().Create().WithName(name).WithVote(vote).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
