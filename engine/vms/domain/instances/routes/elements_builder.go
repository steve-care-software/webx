package routes

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type elementsBuilder struct {
	hashAdapter hash.Adapter
	list        []Element
}

func createElementsBuilder(
	hashAdapter hash.Adapter,
) ElementsBuilder {
	out := elementsBuilder{
		hashAdapter: hashAdapter,
		list:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementsBuilder) Create() ElementsBuilder {
	return createElementsBuilder(
		app.hashAdapter,
	)
}

// WithList adds a list to the builder
func (app *elementsBuilder) WithList(list []Element) ElementsBuilder {
	app.list = list
	return app
}

// Now builds a new Elements instance
func (app *elementsBuilder) Now() (Elements, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Element in order to build an Elements instance")
	}

	data := [][]byte{}
	for _, oneIns := range app.list {
		data = append(data, oneIns.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createElements(*pHash, app.list), nil
}
