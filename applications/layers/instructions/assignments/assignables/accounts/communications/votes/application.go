package votes

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/communications/votes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	stacks_accounts "github.com/steve-care-software/datastencil/domain/stacks/accounts"
)

type application struct {
	accountBuilder    stacks_accounts.Builder
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	accountBuilder stacks_accounts.Builder,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		accountBuilder:    accountBuilder,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable votes.Vote) (stacks.Assignable, *uint, error) {
	messageVar := assignable.Message()
	message, err := frame.FetchBytes(messageVar)
	if err != nil {
		return nil, nil, err
	}

	ringVar := assignable.Ring()
	ring, err := frame.FetchRing(ringVar)
	if err != nil {
		return nil, nil, err
	}

	accountVar := assignable.Account()
	account, err := frame.FetchAccount(accountVar)
	if err != nil {
		return nil, nil, err
	}

	vote, err := account.Signer().Vote(string(message), ring)
	if err != nil {
		return nil, nil, err
	}

	retStackAccount, err := app.accountBuilder.Create().
		WithVote(vote).
		Now()

	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().
		WithAccount(retStackAccount).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil

}
