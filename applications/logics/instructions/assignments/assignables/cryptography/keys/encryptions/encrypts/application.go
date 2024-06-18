package encrypts

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
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
func (app *application) Execute(frame stacks.Frame, assignable encrypts.Encrypt) (stacks.Assignable, *uint, error) {
	msgVar := assignable.Message()
	msg, err := frame.FetchBytes(msgVar)
	if err != nil {
		code := failures.CouldNotFetchMessageFromFrame
		return nil, &code, err
	}

	pubKeyName := assignable.PublicKey()
	pubKey, err := frame.FetchEncryptorPubKey(pubKeyName)
	if err != nil {
		code := failures.CouldNotFetchEncryptionPublicKeyFromFrame
		return nil, &code, err
	}

	cipher, err := pubKey.Encrypt(msg)
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
