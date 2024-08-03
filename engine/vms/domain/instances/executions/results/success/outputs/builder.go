package outputs

import (
	"errors"

	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	input       []byte
	execute     []byte
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		input:       nil,
		execute:     nil,
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
func (app *builder) WithInput(input []byte) Builder {
	app.input = input
	return app
}

// WithExecute adds an execute to the builder
func (app *builder) WithExecute(execute []byte) Builder {
	app.execute = execute
	return app
}

// Now builds a new Output instance
func (app *builder) Now() (Output, error) {
	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.execute != nil && len(app.execute) <= 0 {
		app.execute = nil
	}

	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to buld an Output instance")
	}

	data := [][]byte{
		app.input,
	}

	if app.execute != nil {
		data = append(data, app.execute)
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.execute != nil {
		return createOutputWithExecute(*pHash, app.input, app.execute), nil
	}

	return createOutput(*pHash, app.input), nil
}
