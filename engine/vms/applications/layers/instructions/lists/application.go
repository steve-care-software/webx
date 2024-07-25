package lists

import (
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/lists/deletes"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/lists/inserts"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/lists"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

type application struct {
	insertApp inserts.Application
	deleteApp deletes.Application
}

func createApplication(
	insertApp inserts.Application,
	deleteApp deletes.Application,
) Application {
	out := application{
		insertApp: insertApp,
		deleteApp: deleteApp,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, list lists.List) (stacks.Assignment, *uint, error) {
	if list.IsInsert() {
		insert := list.Insert()
		return app.insertApp.Execute(frame, insert)
	}

	delete := list.Delete()
	return app.deleteApp.Execute(frame, delete)
}
