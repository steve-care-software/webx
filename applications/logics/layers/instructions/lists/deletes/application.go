package deletes

import (
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/layers/instructions/lists/deletes"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
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
func (app *application) Execute(frame stacks.Frame, assignable deletes.Delete) (stacks.Assignment, *uint, error) {
	indexVar := assignable.Index()
	pIndex, err := frame.FetchUnsignedInt(indexVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return nil, &code, nil
	}

	listVar := assignable.List()
	listAssignables, err := frame.FetchList(listVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, nil
	}

	index := *pIndex
	list := listAssignables.List()
	amount := uint(len(list))
	if index >= amount {
		code := failures.CouldNotFetchElementFromList
		return nil, &code, nil
	}

	topDelimiter := index + amount
	if topDelimiter >= amount {
		code := failures.CouldNotFetchElementFromList
		return nil, &code, nil
	}

	remainingList := append(list[0:index], list[index+1:topDelimiter]...)
	newAssignables, err := app.assignablesBuilder.Create().WithList(remainingList).Now()
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
