package keys

import (
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	"github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	encApp encryptions.Application
	sigApp signatures.Application
}

func createApplication(
	encApp encryptions.Application,
	sigApp signatures.Application,
) Application {
	out := application{
		encApp: encApp,
		sigApp: sigApp,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable keys.Key) (stacks.Assignable, *uint, error) {
	if assignable.IsEncryption() {
		enc := assignable.Encryption()
		return app.encApp.Execute(frame, enc)
	}

	signature := assignable.Signature()
	return app.sigApp.Execute(frame, signature)
}
