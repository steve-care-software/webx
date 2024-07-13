package executions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/executions/layers"
	"github.com/steve-care-software/historydb/domain/databases"
	"github.com/steve-care-software/historydb/domain/hash"
)

type executionBuilder struct {
	hashAdapter hash.Adapter
	layer       layers.Layer
	database    databases.Database
}

func createExecutionBuilder(
	hashAdapter hash.Adapter,
) ExecutionBuilder {
	out := executionBuilder{
		hashAdapter: hashAdapter,
		layer:       nil,
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

// WithLayer adds a layer to the executionBuilder
func (app *executionBuilder) WithLayer(layer layers.Layer) ExecutionBuilder {
	app.layer = layer
	return app
}

// WithDatabase adds database to the executionBuilder
func (app *executionBuilder) WithDatabase(database databases.Database) ExecutionBuilder {
	app.database = database
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.layer == nil {
		return nil, errors.New("the layer is mandatory in order to build an Execution instance")
	}

	if app.database == nil {
		return nil, errors.New("the database is mandatory in order to build an Execution instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.layer.Hash().Bytes(),
		app.database.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createExecution(*pHash, app.layer, app.database), nil
}
