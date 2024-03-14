package decrypts

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/encryptors"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	encryptor         encryptors.Encryptor
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	encryptor encryptors.Encryptor,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		encryptor:         encryptor,
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

	passVar := assignable.Password()
	password, err := frame.FetchBytes(passVar)
	if err != nil {
		code := failures.CouldNotFetchPasswordFromFrame
		return nil, &code, err
	}

	result, err := app.encryptor.Decrypt(cipher, password)
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
