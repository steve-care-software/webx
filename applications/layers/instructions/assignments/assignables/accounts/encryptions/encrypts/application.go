package encrypts

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
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
func (app *application) Execute(frame stacks.Frame, assignable encrypts.Encrypt) (stacks.Assignable, *uint, error) {
	msgVar := assignable.Message()
	message, err := frame.FetchBytes(msgVar)
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

	cipher, err := account.Encryptor().Public().Encrypt(message)
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().
		WithBytes(cipher).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
