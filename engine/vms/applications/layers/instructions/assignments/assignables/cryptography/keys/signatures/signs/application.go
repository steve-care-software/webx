package signs

import (
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

type application struct {
	createApp   creates.Application
	validateApp validates.Application
}

func createApplication(
	createApp creates.Application,
	validateApp validates.Application,
) Application {
	out := application{
		createApp:   createApp,
		validateApp: validateApp,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable signs.Sign) (stacks.Assignable, *uint, error) {
	if assignable.IsCreate() {
		create := assignable.Create()
		return app.createApp.Execute(frame, create)
	}

	validate := assignable.Validate()
	return app.validateApp.Execute(frame, validate)
}
