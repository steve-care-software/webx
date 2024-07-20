package votes

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	"github.com/steve-care-software/datastencil/domain/stacks"
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
func (app *application) Execute(frame stacks.Frame, assignable votes.Vote) (stacks.Assignable, *uint, error) {
	if assignable.IsCreate() {
		create := assignable.Create()
		return app.createApp.Execute(frame, create)
	}

	validate := assignable.Validate()
	return app.validateApp.Execute(frame, validate)
}
