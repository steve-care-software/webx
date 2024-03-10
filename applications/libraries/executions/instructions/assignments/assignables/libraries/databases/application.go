package databases

import (
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/assignments/assignables/libraries/databases/repositories"
	"github.com/steve-care-software/datastencil/applications/libraries/executions/instructions/assignments/assignables/libraries/databases/services"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/libraries/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	execRepositoryApp repositories.Application
	execServiceApp    services.Application
}

func createApplication(
	execRepositoryApp repositories.Application,
	execServiceApp services.Application,
) Application {
	out := application{
		execRepositoryApp: execRepositoryApp,
		execServiceApp:    execServiceApp,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable databases.Database) (stacks.Assignable, error) {
	if assignable.IsRepository() {
		repository := assignable.Repository()
		return app.execRepositoryApp.Execute(frame, repository)
	}

	service := assignable.Service()
	return app.execServiceApp.Execute(frame, service)
}
