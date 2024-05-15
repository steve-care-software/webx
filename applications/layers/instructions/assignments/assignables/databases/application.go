package databases

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/commits"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/deletes"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/modifications"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/databases/retrieves"
	assignables_databases "github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	actionApp       actions.Application
	commitApp       commits.Application
	databaseApp     databases.Application
	deleteApp       deletes.Application
	modificationApp modifications.Application
	retrieveApp     retrieves.Application
}

func createApplication(
	actionApp actions.Application,
	commitApp commits.Application,
	databaseApp databases.Application,
	deleteApp deletes.Application,
	modificationApp modifications.Application,
	retrieveApp retrieves.Application,
) Application {
	out := application{
		actionApp:       actionApp,
		commitApp:       commitApp,
		databaseApp:     databaseApp,
		deleteApp:       deleteApp,
		modificationApp: modificationApp,
		retrieveApp:     retrieveApp,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable assignables_databases.Database) (stacks.Assignable, *uint, error) {
	if assignable.IsAction() {
		action := assignable.Action()
		return app.actionApp.Execute(frame, action)
	}

	if assignable.IsCommit() {
		commit := assignable.Commit()
		return app.commitApp.Execute(frame, commit)
	}

	if assignable.IsDatabase() {
		database := assignable.Database()
		return app.databaseApp.Execute(frame, database)
	}

	if assignable.IsDelete() {
		delete := assignable.Delete()
		return app.deleteApp.Execute(frame, delete)
	}

	if assignable.IsModification() {
		modifications := assignable.Modification()
		return app.modificationApp.Execute(frame, modifications)
	}

	retrieve := assignable.Retrieve()
	return app.retrieveApp.Execute(frame, retrieve)
}
