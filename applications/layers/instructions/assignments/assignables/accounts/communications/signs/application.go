package signs

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/accounts/communications/signs"
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
func (app *application) Execute(frame stacks.Frame, assignable signs.Sign) (stacks.Assignable, *uint, error) {
	messageVar := assignable.Message()
	message, err := frame.FetchBytes(messageVar)
	if err != nil {
		code := failures.CouldNotFetchMessageFromFrame
		return nil, &code, err
	}

	accountVar := assignable.Account()
	account, err := frame.FetchAccount(accountVar)
	if err != nil {
		code := failures.CouldNotFetchAccountFromFrame
		return nil, &code, err
	}

	sig, err := account.Signer().Sign(string(message))
	if err != nil {
		return nil, nil, err
	}

	retStackAccount, err := app.accountBuilder.Create().
		WithSignature(sig).
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
