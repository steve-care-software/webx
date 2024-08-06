package merges

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	base        string
	top         string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		base:        "",
		top:         "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithBase adds a base to the builder
func (app *builder) WithBase(base string) Builder {
	app.base = base
	return app
}

// WithTop adds a top to the builder
func (app *builder) WithTop(top string) Builder {
	app.top = top
	return app
}

// Now builds a new Merge instance
func (app *builder) Now() (Merge, error) {
	if app.base == "" {
		return nil, errors.New("the base is mandatory in order to build a Merge instance")
	}

	if app.top == "" {
		return nil, errors.New("the top is mandatory in order to build a Merge instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.base),
		[]byte(app.top),
	})

	if err != nil {
		return nil, err
	}

	return createMerge(*pHash, app.base, app.top), nil
}
