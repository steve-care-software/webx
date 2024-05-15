package actions

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	database_actions "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications"

	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/domain/stacks"
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
		return nil, nil, err
	}

	modificationsList := []modifications.Modification{}
	modifList := modifAssignables.List()
	for _, oneAssignable := range modifList {
		if !oneAssignable.IsModification() {
			code := failures.CouldNotFetchModificationFromList
			return nil, &code, nil
		}

		modificationsList = append(modificationsList, oneAssignable.Modification())
	}

	modifications, err := app.modificationsBuilder.Create().WithList(modificationsList).Now()
	if err != nil {
		return nil, nil, err
	}

	pathVar := assignable.Path()
	pathAssignables, err := frame.FetchList(pathVar)
	if err != nil {
		return nil, nil, err
	}

	pathValues := []string{}
	pathList := pathAssignables.List()
	for _, oneAssignable := range pathList {
		if !oneAssignable.IsString() {
			code := failures.CouldNotFetchStringFromList
			return nil, &code, nil
		}

		pathValues = append(pathValues, *oneAssignable.String())
	}

	action, err := app.actionBuilder.Create().WithModifications(modifications).WithPath(pathValues).Now()
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().WithAction(action).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
