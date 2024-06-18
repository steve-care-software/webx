package deletes

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/lists/deletes"
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
		return nil, &code, err
	}

	listVar := assignable.List()
	listAssignables, err := frame.FetchList(listVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, err
	}

	index := *pIndex
	list := listAssignables.List()
	amount := uint(len(list))
	if index >= amount {
		code := failures.CouldNotFetchElementFromList
		str := fmt.Sprintf("the index (%d) exceeds the top delimiter (%d) of the provided list (%s)", index, amount-1, listVar)
		return nil, &code, errors.New(str)
	}

	remainingList := append(list[0:index], list[index+1:]...)
	newAssignables, err := app.assignablesBuilder.Create().
		WithList(remainingList).
		Now()

	if err != nil {
		return nil, nil, err
	}

	newAssignable, err := app.assignableBuiler.Create().
		WithList(newAssignables).
		Now()

	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignmentBuilder.Create().
		WithAssignable(newAssignable).
		WithName(listVar).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
