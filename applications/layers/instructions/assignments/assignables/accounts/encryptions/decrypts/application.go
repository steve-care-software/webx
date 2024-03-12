package decrypts

import (
	"github.com/steve-care-software/datastencil/domain/instances/libraries/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	assignableBuilder stacks.AssignableBuilder
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable decrypts.Decrypt) (stacks.Assignable, *uint, error) {
	cipherVar := assignable.Cipher()
	cipher, err := frame.FetchBytes(cipherVar)
	if err != nil {
		return nil, nil, err
	}

	accountVar := assignable.Account()
	account, err := frame.FetchAccount(accountVar)
	if err != nil {
		return nil, nil, err
	}

	result, err := account.Encryptor().Decrypt(cipher)
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().
		WithBytes(result).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
