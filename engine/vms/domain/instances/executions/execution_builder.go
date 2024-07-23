package executions

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results"
	source_layers "github.com/steve-care-software/webx/engine/vms/domain/instances/layers"
)

type executionBuilder struct {
	hashAdapter hash.Adapter
	input       []byte
	source      source_layers.Layer
	result      results.Result
}

func createExecutionBuilder(
	hashAdapter hash.Adapter,
) ExecutionBuilder {
	out := executionBuilder{
		hashAdapter: hashAdapter,
		input:       nil,
		source:      nil,
		result:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder(
		app.hashAdapter,
	)
}

// WithInput adds an input to the builder
func (app *executionBuilder) WithInput(input []byte) ExecutionBuilder {
	app.input = input
	return app
}

// WithSource adds a source to the builder
func (app *executionBuilder) WithSource(source source_layers.Layer) ExecutionBuilder {
	app.source = source
	return app
}

// WithResult adds a result to the builder
func (app *executionBuilder) WithResult(result results.Result) ExecutionBuilder {
	app.result = result
	return app
}

// Now builds a new Layer instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.source == nil {
		return nil, errors.New("the source is mandatory in order to build an Execution instance")
	}

	if app.result == nil {
		return nil, errors.New("the result is mandatory in order to build an Execution instance")
	}

	data := [][]byte{
		app.source.Hash().Bytes(),
		app.result.Hash().Bytes(),
	}

	if app.input != nil {
		data = append(data, app.input)
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.input != nil {
		return createExecutionWithInput(*pHash, app.source, app.result, app.input), nil
	}

	return createExecution(*pHash, app.source, app.result), nil
}
