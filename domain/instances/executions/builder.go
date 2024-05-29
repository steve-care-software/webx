package executions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/executions/links"
)

type builder struct {
	hashAdapter hash.Adapter
	logic       links.Link
	database    databases.Database
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		logic:       nil,
		database:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithLogic adds logic to the builder
func (app *builder) WithLogic(logic links.Link) Builder {
	app.logic = logic
	return app
}

// WithDatabase adds database to the builder
func (app *builder) WithDatabase(database databases.Database) Builder {
	app.database = database
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
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
