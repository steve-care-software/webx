package inserts

import (
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/lists/inserts"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks/failures"
)

type application struct {
	assignableBuiler   stacks.AssignableBuilder
	assignablesBuilder stacks.AssignablesBuilder
	assignmentBuilder  stacks.AssignmentBuilder
}

func createApplication(
	assignableBuiler stacks.AssignableBuilder,
	assignablesBuilder stacks.AssignablesBuilder,
	assignmentBuilder stacks.AssignmentBuilder,
) Application {
	out := application{
		assignableBuiler:   assignableBuiler,
		assignablesBuilder: assignablesBuilder,
		assignmentBuilder:  assignmentBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable inserts.Insert) (stacks.Assignment, *uint, error) {
	listVar := assignable.List()
	listAssignables, err := frame.FetchList(listVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, err
	}

	elementVar := assignable.Element()
	fetched, err := frame.Fetch(elementVar)
	if err != nil {
		code := failures.CouldNotFetchFromFrame
		return nil, &code, err
	}

	list := listAssignables.List()
	list = append(list, fetched)
	newAssignables, err := app.assignablesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	newAssignable, err := app.assignableBuiler.Create().WithList(newAssignables).Now()
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignmentBuilder.Create().WithAssignable(newAssignable).WithName(listVar).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
