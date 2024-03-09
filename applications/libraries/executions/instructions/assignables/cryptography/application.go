package cryptography

import (
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execDecryptApp decrypts.Application
	execEncryptApp encrypts.Application
}

func createApplication(
	execDecryptApp decrypts.Application,
	execEncryptApp encrypts.Application,
) Application {
	out := application{
		execDecryptApp: execDecryptApp,
		execEncryptApp: execEncryptApp,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable cryptography.Cryptography) (stacks.Assignable, error) {
	if assignable.IsDecrypt() {
		decrypt := assignable.Decrypt()
		return app.execDecryptApp.Execute(frame, decrypt)
	}

	encrypt := assignable.Encrypt()
	return app.execEncryptApp.Execute(frame, encrypt)
}
