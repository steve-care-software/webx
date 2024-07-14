package executes

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes/inputs"
	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	context     string
	input       inputs.Input
	ret         string
	layer       inputs.Input
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		context:     "",
		input:       nil,
		ret:         "",
		layer:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithContext adds a context to the builder
func (app *builder) WithContext(context string) Builder {
	app.context = context
	return app
}

// WithInput adds an input to the builder
func (app *builder) WithInput(input inputs.Input) Builder {
	app.input = input
	return app
}

// WithReturn adds a return to the builder
func (app *builder) WithReturn(ret string) Builder {
	app.ret = ret
	return app
}

// WithLayer adds a layer to the builder
func (app *builder) WithLayer(layer inputs.Input) Builder {
	app.layer = layer
	return app
}

// Now builds a new Execute instance
func (app *builder) Now() (Execute, error) {
	if app.context == "" {
		return nil, errors.New("the context is mandatory in order to build an Execute instance")
	}

	if app.input == nil {
		return nil, errors.New("the input is mandatory in order to build an Execute instance")
	}

	if app.ret == "" {
		return nil, errors.New("the return is mandatory in order to build an Execute instance")
	}

	bytes := [][]byte{
		[]byte(app.context),
		app.input.Hash().Bytes(),
		[]byte(app.ret),
	}

	if app.layer != nil {
		bytes = append(bytes, app.layer.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(bytes)
	if err != nil {
		return nil, err
	}

	if app.layer != nil {
		return createExecuteWithLayer(*pHash, app.context, app.input, app.ret, app.layer), nil
	}

	return createExecute(*pHash, app.context, app.input, app.ret), nil
}
