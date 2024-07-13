package outputs

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/layers/outputs/kinds"
	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	variable    string
	kind        kinds.Kind
	execute     []string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		variable:    "",
		kind:        nil,
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

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// WithKind adds a kind to the builder
func (app *builder) WithKind(kind kinds.Kind) Builder {
	app.kind = kind
	return app
}

// WithExecute adds an execute to the builder
func (app *builder) WithExecute(execute []string) Builder {
	app.execute = execute
	return app
}

// Now builds a new Output instance
func (app *builder) Now() (Output, error) {
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

	if app.execute != nil && len(app.execute) <= 0 {
		app.execute = nil
	}

	if app.execute != nil {
		for _, oneArg := range app.execute {
			data = append(data, []byte(oneArg))
		}
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.execute != nil {
		return createOutputWithExecute(*pHash, app.variable, app.kind, app.execute), nil
	}

	return createOutput(*pHash, app.variable, app.kind), nil
}
