package instructions

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Instruction
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
func (app *builder) WithList(list []Instruction) Builder {
	app.list = list
	return app
}

// Now builds a new Instructions instance
func (app *builder) Now() (Instructions, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Instruction in order to build an Instructions instance")
	}

	data := [][]byte{}
	for _, oneIns := range app.list {
		data = append(data, oneIns.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createInstructions(*pHash, app.list), nil
}
