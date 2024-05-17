package inserts

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/lists/inserts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
}

func createApplication() Application {
	out := application{}
	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable inserts.Insert) ([]stacks.Assignable, *uint, error) {
	listVar := assignable.List()
	listAssignables, err := frame.FetchList(listVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, nil
	}

	elementVar := assignable.Element()
	newAssignable, err := frame.Fetch(elementVar)
	if err != nil {
		code := failures.CouldNotFetchFromFrame
		return nil, &code, nil
	}

	list := listAssignables.List()
	list = append(list, newAssignable)
	return list, nil, nil
}
