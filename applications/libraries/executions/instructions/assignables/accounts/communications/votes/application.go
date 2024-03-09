package votes

import (
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/communications/votes"
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
func (app *application) Execute(frame stacks.Frame, assignable votes.Vote) (stacks.Assignable, error) {
	messageVar := assignable.Message()
	message, err := frame.FetchBytes(messageVar)
	if err != nil {
		return nil, err
	}

	ringVar := assignable.Ring()
	ring, err := frame.FetchRing(ringVar)
	if err != nil {
		return nil, err
	}

	accountVar := assignable.Account()
	account, err := frame.FetchAccount(accountVar)
	if err != nil {
		return nil, err
	}

	vote, err := account.Signer().Vote(message, ring)
	if err != nil {
		return nil, err
	}

	stackAccount, err := app.accountBuilder.Create().
		WithVote(vote).
		Now()

	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithAccount(stackAccount).
		Now()

}
