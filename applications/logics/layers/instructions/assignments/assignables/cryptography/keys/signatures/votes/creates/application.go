package creates

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable creates.Create) (stacks.Assignable, *uint, error) {
	msgVar := assignable.Message()
	msg, err := frame.FetchBytes(msgVar)
	if err != nil {
		code := failures.CouldNotFetchMessageFromFrame
		return nil, &code, err
	}

	ringVar := assignable.Ring()
	ring, err := frame.FetchRing(ringVar)
	if err != nil {
		code := failures.CouldNotFetchRingFromFrame
		return nil, &code, err
	}

	pkName := assignable.PrivateKey()
	pk, err := frame.FetchSigner(pkName)
	if err != nil {
		code := failures.CouldNotFetchSignerPrivateKeyFromFrame
		return nil, &code, err
	}

	vote, err := pk.Vote(string(msg), ring)
	if err != nil {
		code := failures.CouldNotVoteOnMessageInFrame
		return nil, &code, err
	}

	ins, err := app.assignableBuilder.Create().WithVote(vote).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
