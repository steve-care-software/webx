package encryptions

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions/decrypts"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/accounts/encryptions/encrypts"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts/encryptions"
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
func (app *application) Execute(frame stacks.Frame, assignable encryptions.Encryption) (stacks.Assignable, error) {
	if assignable.IsEncrypt() {
		encrypt := assignable.Encrypt()
		return app.execEncryptApp.Execute(frame, encrypt)
	}

	decrypt := assignable.Decrypt()
	return app.execDecryptApp.Execute(frame, decrypt)
}
