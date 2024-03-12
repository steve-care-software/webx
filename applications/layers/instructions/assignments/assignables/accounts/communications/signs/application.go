package signs

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/communications/signs"
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
func (app *application) Execute(frame stacks.Frame, assignable signs.Sign) (stacks.Assignable, error) {
	messageVar := assignable.Message()
	message, err := frame.FetchBytes(messageVar)
	if err != nil {
		return nil, err
	}

	accountVar := assignable.Account()
	account, err := frame.FetchAccount(accountVar)
	if err != nil {
		return nil, err
	}

	sig, err := account.Signer().Sign(string(message))
	if err != nil {
		return nil, err
	}

	retStackAccount, err := app.accountBuilder.Create().
		WithSignature(sig).
		Now()

	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithAccount(retStackAccount).
		Now()

}
