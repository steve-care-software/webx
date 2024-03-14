package decrypts

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
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
func (app *application) Execute(frame stacks.Frame, assignable decrypts.Decrypt) (stacks.Assignable, *uint, error) {
	cipherVar := assignable.Cipher()
	cipher, err := frame.FetchBytes(cipherVar)
	if err != nil {
		code := failures.CouldNotFetchCipherFromFrame
		return nil, &code, err
	}

	accountVar := assignable.Account()
	account, err := frame.FetchAccount(accountVar)
	if err != nil {
		code := failures.CouldNotFetchAccountFromFrame
		return nil, &code, err
	}

	result, err := account.Encryptor().Decrypt(cipher)
	if err != nil {
		code := failures.CouldNotDecryptCipher
		return nil, &code, err
	}

	ins, err := app.assignableBuilder.Create().
		WithBytes(result).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
