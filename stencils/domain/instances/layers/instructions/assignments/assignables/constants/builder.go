package constants

import (
	"errors"

	"github.com/steve-care-software/datastencil/states/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        []Constant
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
func (app *builder) WithList(list []Constant) Builder {
	app.list = list
	return app
}

// Now builds a new Constants instance
func (app *builder) Now() (Constants, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Constant in order to build a Constants instance")
	}

	data := [][]byte{}
	for _, oneConstant := range app.list {
		data = append(data, oneConstant.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createConstants(*pHash, app.list), nil
}
