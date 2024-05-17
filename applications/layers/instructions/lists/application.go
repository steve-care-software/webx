package lists

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/lists/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/lists"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	insertApp          inserts.Application
	deleteApp          deletes.Application
	assignablesBuilder stacks.AssignablesBuilder
}

func createApplication(
	insertApp inserts.Application,
	deleteApp deletes.Application,
	assignablesBuilder stacks.AssignablesBuilder,
) Application {
	out := application{
		insertApp:          insertApp,
		deleteApp:          deleteApp,
		assignablesBuilder: assignablesBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, list lists.List) (stacks.Assignables, *uint, error) {
	builder := app.assignablesBuilder.Create()
	if list.IsInsert() {
		insert := list.Insert()
		retList, code, err := app.insertApp.Execute(frame, insert)
		if err != nil {
			return nil, code, err
		}

		builder.WithList(retList)
	}

	if list.IsDelete() {
		delete := list.Delete()
		retList, code, err := app.deleteApp.Execute(frame, delete)
		if err != nil {
			return nil, code, err
		}

		builder.WithList(retList)
	}

	retIns, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return retIns, nil, nil
}
