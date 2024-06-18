package cryptography

import (
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/applications/logics/instructions/assignments/assignables/cryptography/keys"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execDecryptApp decrypts.Application
	execEncryptApp encrypts.Application
	keyApp         keys.Application
}

func createApplication(
	execDecryptApp decrypts.Application,
	execEncryptApp encrypts.Application,
	keyApp keys.Application,
) Application {
	out := application{
		execDecryptApp: execDecryptApp,
		execEncryptApp: execEncryptApp,
		keyApp:         keyApp,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable cryptography.Cryptography) (stacks.Assignable, *uint, error) {
	if assignable.IsDecrypt() {
		decrypt := assignable.Decrypt()
		return app.execDecryptApp.Execute(frame, decrypt)
	}

	if assignable.IsEncrypt() {
		encrypt := assignable.Encrypt()
		return app.execEncryptApp.Execute(frame, encrypt)
	}

	key := assignable.Key()
	return app.keyApp.Execute(frame, key)
}
