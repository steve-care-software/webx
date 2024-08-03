package executions

import (
	"errors"

	"github.com/steve-care-software/webx/engine/hashes/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Execution
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Execution) Builder {
	app.list = list
	return app
}

// Now builds a new Executions instance
func (app *builder) Now() (Executions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Execution in order to build an Executions instance")
	}

	bytes := [][]byte{}
	for _, oneIns := range app.list {
		bytes = append(bytes, oneIns.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(bytes)
	if err != nil {
		return nil, err
	}

	return createExecutions(*pHash, app.list), nil
}
