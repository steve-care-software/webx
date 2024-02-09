package layers

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type outputBuilder struct {
	hashAdapter hash.Adapter
	variable    string
	kind        Kind
	execute     string
}

func createOutputBuilder(
	hashAdapter hash.Adapter,
) OutputBuilder {
	out := outputBuilder{
		hashAdapter: hashAdapter,
		variable:    "",
		kind:        nil,
		execute:     "",
	}

	return &out
}

// Create initializes the builder
func (app *outputBuilder) Create() OutputBuilder {
	return createOutputBuilder(
		app.hashAdapter,
	)
}

// WithVariable adds a variable to the builder
func (app *outputBuilder) WithVariable(variable string) OutputBuilder {
	app.variable = variable
	return app
}

// WithKind adds a kind to the builder
func (app *outputBuilder) WithKind(kind Kind) OutputBuilder {
	app.kind = kind
	return app
}

// WithExecute adds an execute to the builder
func (app *outputBuilder) WithExecute(execute string) OutputBuilder {
	app.execute = execute
	return app
}

// Now builds a new Output instance
func (app *outputBuilder) Now() (Output, error) {
	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build an Output instance")
	}

	if app.kind == nil {
		return nil, errors.New("the kind is mandatory in order to build an Output instance")
	}

	data := [][]byte{
		[]byte(app.variable),
		app.kind.Hash().Bytes(),
	}

	if app.execute != "" {
		data = append(data, []byte(app.execute))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.execute != "" {
		return createOutputWithExecute(*pHash, app.variable, app.kind, app.execute), nil
	}

	return createOutput(*pHash, app.variable, app.kind), nil
}
