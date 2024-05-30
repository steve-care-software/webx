package executions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links"
)

type executionBuilder struct {
	hashAdapter hash.Adapter
	logic       links.Link
	database    databases.Database
}

func createExecutionBuilder(
	hashAdapter hash.Adapter,
) ExecutionBuilder {
	out := executionBuilder{
		hashAdapter: hashAdapter,
		logic:       nil,
		database:    nil,
	}

	return &out
}

// Create initializes the executionBuilder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder(
		app.hashAdapter,
	)
}

// WithLogic adds logic to the executionBuilder
func (app *executionBuilder) WithLogic(logic links.Link) ExecutionBuilder {
	app.logic = logic
	return app
}

// WithDatabase adds database to the executionBuilder
func (app *executionBuilder) WithDatabase(database databases.Database) ExecutionBuilder {
	app.database = database
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.logic == nil {
		return nil, errors.New("the logic is mandatory in order to build an Execution instance")
	}

	if app.database == nil {
		return nil, errors.New("the database is mandatory in order to build an Execution instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.logic.Hash().Bytes(),
		app.database.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createExecution(*pHash, app.logic, app.database), nil
}
