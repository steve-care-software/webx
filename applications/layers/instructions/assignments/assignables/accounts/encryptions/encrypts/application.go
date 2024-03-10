package encrypts

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	assignableBuilder stacks.AssignableBuilder
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable encrypts.Encrypt) (stacks.Assignable, error) {
	msgVar := assignable.Message()
	message, err := frame.FetchBytes(msgVar)
	if err != nil {
		return nil, err
	}

	accountVar := assignable.Account()
	account, err := frame.FetchAccount(accountVar)
	if err != nil {
		return nil, err
	}

	cipher, err := account.Encryptor().Public().Encrypt(message)
	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithBytes(cipher).
		Now()
}
