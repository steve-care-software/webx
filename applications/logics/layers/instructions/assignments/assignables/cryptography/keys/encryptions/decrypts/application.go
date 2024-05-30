package decrypts

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
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
func (app *application) Execute(frame stacks.Frame, assignable decrypts.Decrypt) (stacks.Assignable, *uint, error) {
	cipherVar := assignable.Cipher()
	cipher, err := frame.FetchBytes(cipherVar)
	if err != nil {
		code := failures.CouldNotFetchCipherFromFrame
		return nil, &code, err
	}

	pkName := assignable.PrivateKey()
	pk, err := frame.FetchEncryptor(pkName)
	if err != nil {
		code := failures.CouldNotFetchEncryptionPrivateKeyFromFrame
		return nil, &code, err
	}

	value, err := pk.Decrypt(cipher)
	if err != nil {
		code := failures.CouldNotDecryptCipher
		return nil, &code, err
	}

	ins, err := app.assignableBuilder.Create().
		WithBytes(value).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
