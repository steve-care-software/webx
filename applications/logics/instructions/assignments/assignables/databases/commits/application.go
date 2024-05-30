package commits

import (
	databases_commits "github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/commits"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	assignableBuilder stacks.AssignableBuilder
	actionsBuilder    actions.Builder
	commitBuilder     databases_commits.Builder
}

func createApplication(
	assignableBuilder stacks.AssignableBuilder,
	actionsBuilder actions.Builder,
	commitBuilder databases_commits.Builder,
) Application {
	out := application{
		assignableBuilder: assignableBuilder,
		actionsBuilder:    actionsBuilder,
		commitBuilder:     commitBuilder,
	}

	return &out
}

// Execute executes an application
func (app *application) Execute(frame stacks.Frame, assignable commits.Commit) (stacks.Assignable, *uint, error) {
	actionsVar := assignable.Actions()
	actionsAssignables, err := frame.FetchList(actionsVar)
	if err != nil {
		return nil, nil, err
	}

	actionList := []actions.Action{}
	actionAssList := actionsAssignables.List()
	for _, oneAssignable := range actionAssList {
		if !oneAssignable.IsModification() {
			code := failures.CouldNotFetchActionFromList
			return nil, &code, nil
		}

		actionList = append(actionList, oneAssignable.Action())
	}

	actions, err := app.actionsBuilder.Create().WithList(actionList).Now()
	if err != nil {
		return nil, nil, err
	}

	descriptionVar := assignable.Description()
	description, err := frame.FetchString(descriptionVar)
	if err != nil {
		code := failures.CouldNotFetchStringFromFrame
		return nil, &code, nil
	}

	builder := app.commitBuilder.Create().WithActions(actions).WithDescription(description)
	if assignable.HashParent() {
		parentVar := assignable.Parent()
		parentHash, err := frame.FetchHash(parentVar)
		if err != nil {
			code := failures.CouldNotFetchHashVariableFromFrame
			return nil, &code, err
		}

		builder.WithParent(parentHash)
	}

	commit, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().WithCommit(commit).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
