package databases

import (
	databases_databases "github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	assignableBuilder stacks.AssignableBuilder
	databaseBuilder   databases_databases.Builder
}

func createApplication(
	assignableBuilder stacks.AssignableBuilder,
	databaseBuilder databases_databases.Builder,
) Application {
	out := application{
		assignableBuilder: assignableBuilder,
		databaseBuilder:   databaseBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable databases.Database) (stacks.Assignable, *uint, error) {
	pathVar := assignable.Path()
	pathAssignables, err := frame.FetchList(pathVar)
	if err != nil {
		code := failures.CouldNotFetchListFromFrame
		return nil, &code, nil
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

	descriptionVar := assignable.Description()
	description, err := frame.FetchString(descriptionVar)
	if err != nil {
		code := failures.CouldNotFetchStringFromFrame
		return nil, &code, nil
	}

	headVar := assignable.Head()
	commit, err := frame.FetchCommit(headVar)
	if err != nil {
		code := failures.CouldNotFetchCommitFromFrame
		return nil, &code, nil
	}

	isActiveVar := assignable.IsActive()
	isActive, err := frame.FetchBool(isActiveVar)
	if err != nil {
		code := failures.CouldNotFetchBoolFromFrame
		return nil, &code, nil
	}

	builder := app.databaseBuilder.Create().WithPath(pathValues).WithDescription(description).WithHead(commit)
	if isActive {
		builder.IsActive()
	}

	database, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().WithDatabase(database).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
