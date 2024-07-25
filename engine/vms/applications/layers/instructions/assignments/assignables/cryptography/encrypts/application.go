package encrypts

import (
	"github.com/steve-care-software/webx/engine/vms/domain/encryptors"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/failures"
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
func (app *application) Execute(frame stacks.Frame, assignable encrypts.Encrypt) (stacks.Assignable, *uint, error) {
	msgVar := assignable.Message()
	message, err := frame.FetchBytes(msgVar)
	if err != nil {
		code := failures.CouldNotFetchMessageFromFrame
		return nil, &code, err
	}

	passVar := assignable.Password()
	password, err := frame.FetchBytes(passVar)
	if err != nil {
		code := failures.CouldNotFetchPasswordFromFrame
		return nil, &code, err
	}

	cipher, err := app.encryptor.Encrypt(message, password)
	if err != nil {
		code := failures.CouldNotEncryptMessage
		return nil, &code, err
	}

	ins, err := app.assignableBuilder.Create().
		WithBytes(cipher).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
