package executions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	input       string
	layer       string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		input:       "",
		layer:       "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input string) Builder {
	app.input = input
	return app
}

// WithLayer adds a layer to the builder
func (app *builder) WithLayer(layer string) Builder {
	app.layer = layer
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
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
