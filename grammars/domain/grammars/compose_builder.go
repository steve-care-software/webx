package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type composeBuilder struct {
	hashAdapter hash.Adapter
	list        []ComposeElement
}

func createComposeBuilder(
	hashAdapter hash.Adapter,
) ComposeBuilder {
	out := composeBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *composeBuilder) Create() ComposeBuilder {
	return createComposeBuilder(app.hashAdapter)
}

// WithList adds a list to the builder
func (app *composeBuilder) WithList(list []ComposeElement) ComposeBuilder {
	app.list = list
	return app
}

// Now builds a new Compose instance
func (app *composeBuilder) Now() (Compose, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 ComposeElement in order to build a Compose instance")
	}

	data := [][]byte{}
	for _, oneElement := range app.list {
		data = append(data, oneElement.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createCompose(*pHash, app.list), nil
}
