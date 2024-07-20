package results

import (
	"errors"

	"github.com/steve-care-software/datastencil/states/domain/hash"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/executions/results/interruptions"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/executions/results/success"
)

type builder struct {
	hashAdapter  hash.Adapter
	success      success.Success
	interruption interruptions.Interruption
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		success:      nil,
		interruption: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithSuccess adds a success to the builder
func (app *builder) WithSuccess(success success.Success) Builder {
	app.success = success
	return app
}

// WithInterruption adds an interruption to the builder
func (app *builder) WithInterruption(interruption interruptions.Interruption) Builder {
	app.interruption = interruption
	return app
}

// Now builds a new Result instance
func (app *builder) Now() (Result, error) {
	data := [][]byte{}
	if app.success != nil {
		data = append(data, []byte("isSuccess"))
		data = append(data, app.success.Hash().Bytes())
	}

	if app.interruption != nil {
		data = append(data, []byte("IsInterruption"))
		data = append(data, app.interruption.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Result is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.success != nil {
		return createResultWithSuccess(*pHash, app.success), nil
	}

	return createResultWithInterruption(*pHash, app.interruption), nil
}
