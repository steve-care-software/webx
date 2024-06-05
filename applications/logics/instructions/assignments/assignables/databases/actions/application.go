package actions

import (
	"errors"
	"fmt"

	database_actions "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	assignableBuilder    stacks.AssignableBuilder
	modificationsBuilder modifications.Builder
	actionBuilder        database_actions.ActionBuilder
}

func createApplication(
	assignableBuilder stacks.AssignableBuilder,
	modificationsBuilder modifications.Builder,
	actionBuilder database_actions.ActionBuilder,
) Application {
	out := application{
		assignableBuilder:    assignableBuilder,
		modificationsBuilder: modificationsBuilder,
		actionBuilder:        actionBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable actions.Action) (stacks.Assignable, *uint, error) {
	modifVar := assignable.Modifications()
	modifAssignables, err := frame.FetchList(modifVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, err
	}

	modificationsList := []modifications.Modification{}
	modifList := modifAssignables.List()
	for _, oneAssignable := range modifList {
		if !oneAssignable.IsModification() {
			code := failures.CouldNotFetchModificationFromList
			str := fmt.Sprintf("the list (name: %s) was expected to contain Modification instances", modifVar)
			return nil, &code, errors.New(str)
		}

		modificationsList = append(modificationsList, oneAssignable.Modification())
	}

	modifications, err := app.modificationsBuilder.Create().
		WithList(modificationsList).
		Now()

	if err != nil {
		return nil, nil, err
	}

	pathVar := assignable.Path()
	pathAssignables, err := frame.FetchList(pathVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, err
	}

	pathValues := []string{}
	pathList := pathAssignables.List()
	for _, oneAssignable := range pathList {
		if !oneAssignable.IsString() {
			code := failures.CouldNotFetchStringFromList
			str := fmt.Sprintf("the list (name: %s) was expected to contain string values", modifVar)
			return nil, &code, errors.New(str)
		}

		pathValues = append(pathValues, *oneAssignable.String())
	}

	action, err := app.actionBuilder.Create().
		WithModifications(modifications).
		WithPath(pathValues).
		Now()

	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().
		WithAction(action).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
