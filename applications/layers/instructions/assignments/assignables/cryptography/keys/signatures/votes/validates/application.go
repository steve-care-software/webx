package validates

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	"github.com/steve-care-software/datastencil/domain/keys/signers"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	assignableBuilder stacks.AssignableBuilder
	voteAdapter       signers.VoteAdapter
}

func createApplication(
	assignableBuilder stacks.AssignableBuilder,
	voteAdapter signers.VoteAdapter,
) Application {
	out := application{
		assignableBuilder: assignableBuilder,
		voteAdapter:       voteAdapter,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable validates.Validate) (stacks.Assignable, *uint, error) {
	voteVar := assignable.Vote()
	vote, err := frame.FetchVote(voteVar)
	if err != nil {
		code := failures.CouldNotFetchVoteFromFrame
		return nil, &code, err
	}

	msgVar := assignable.Message()
	msg, err := frame.FetchBytes(msgVar)
	if err != nil {
		code := failures.CouldNotFetchMessageFromFrame
		return nil, &code, err
	}

	ringVar := assignable.HashedRing()
	ringAssignables, err := frame.FetchList(ringVar)
	if err != nil {
		code := failures.CouldNotFetchRingFromFrame
		return nil, &code, err
	}

	hashes := []hash.Hash{}
	list := ringAssignables.List()
	for _, oneAssignable := range list {
		if !oneAssignable.IsHash() {
			code := failures.CouldNotFetchHashFromList
			return nil, &code, nil
		}

		hashes = append(hashes, oneAssignable.Hash())
	}

	validated, err := app.voteAdapter.ToVerification(vote, string(msg), hashes)
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().WithBool(validated).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
