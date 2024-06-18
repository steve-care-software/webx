package inserts

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	list        string
	element     string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		list:        "",
		element:     "",
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
func (app *builder) WithList(list string) Builder {
	app.list = list
	return app
}

// WithElement adds an element to the builder
func (app *builder) WithElement(element string) Builder {
	app.element = element
	return app
}

// Now builds a new Insert
func (app *builder) Now() (Insert, error) {
	if app.list == "" {
		return nil, errors.New("the list is mandatory in order to build an Insert instance")
	}

	if app.element == "" {
		return nil, errors.New("the element is mandatory in order to build an Insert instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.list),
		[]byte(app.element),
	})

	if err != nil {
		return nil, err
	}

	return createInsert(*pHash, app.list, app.element), nil
}
