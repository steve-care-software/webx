package lists

import (
	application_fetches "github.com/steve-care-software/datastencil/stencils/applications/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/lists"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks/failures"
)

type application struct {
	fetchApplication   application_fetches.Application
	assignableBuilder  stacks.AssignableBuilder
	assignablesBuilder stacks.AssignablesBuilder
}

func createApplication(
	fetchApplication application_fetches.Application,
	assignableBuilder stacks.AssignableBuilder,
	assignablesBuilder stacks.AssignablesBuilder,
) Application {
	out := application{
		fetchApplication:   fetchApplication,
		assignableBuilder:  assignableBuilder,
		assignablesBuilder: assignablesBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable lists.List) (stacks.Assignable, *uint, error) {
	if assignable.IsFetch() {
		fetch := assignable.Fetch()
		return app.fetchApplication.Execute(frame, fetch)
	}

	if assignable.IsCreate() {
		elementName := assignable.Create()
		assignable, err := frame.Fetch(elementName)
		if err != nil {
			code := failures.CouldNotFetchFromFrame
			return nil, &code, err
		}

		retAssignables, err := app.assignablesBuilder.Create().WithList([]stacks.Assignable{
			assignable,
		}).Now()
		if err != nil {
			return nil, nil, err
		}

		retAssignable, err := app.assignableBuilder.Create().WithList(retAssignables).Now()
		if err != nil {
			return nil, nil, err
		}

		return retAssignable, nil, nil
	}

	listVar := assignable.Length()
	list, err := frame.FetchList(listVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, err
	}

	length := len(list.List())
	retAssignable, err := app.assignableBuilder.Create().WithUnsignedInt(uint(length)).Now()
	if err != nil {
		return nil, nil, err
	}

	return retAssignable, nil, nil
}
