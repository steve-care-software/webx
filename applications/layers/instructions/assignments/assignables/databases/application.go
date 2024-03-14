package databases

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/repositories"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/services"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/databases"
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
func (app *application) Execute(frame stacks.Frame, assignable databases.Database) (stacks.Assignable, *uint, error) {
	if assignable.IsRepository() {
		repository := assignable.Repository()
		return app.execRepositoryApp.Execute(frame, repository)
	}

	service := assignable.Service()
	return app.execServiceApp.Execute(frame, service)
}
