package layers

import (
	"errors"

	"github.com/steve-care-software/identity/domain/hash"
)

type executionBuilder struct {
	hashAdapter hash.Adapter
	input       string
	layer       string
}

func createExecutionBuilder(
	hashAdapter hash.Adapter,
) ExecutionBuilder {
	out := executionBuilder{
		hashAdapter: hashAdapter,
		input:       "",
		layer:       "",
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
func (app *executionBuilder) WithInput(input string) ExecutionBuilder {
	app.input = input
	return app
}

// WithLayer adds a layer to the builder
func (app *executionBuilder) WithLayer(layer string) ExecutionBuilder {
	app.layer = layer
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.input == "" {
		return nil, errors.New("the input is mandatory in order to build an Execution instance")
	}

	data := [][]byte{
		[]byte(app.input),
	}

	if app.layer != "" {
		data = append(data, []byte(app.layer))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.layer != "" {
		return createExecutionWithLayer(*pHash, app.input, app.layer), nil
	}

	return createExecution(*pHash, app.input), nil
}
