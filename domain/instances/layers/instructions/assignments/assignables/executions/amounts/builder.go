package amounts

import (
	"errors"

	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	context     string
	ret         string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		context:     "",
		ret:         "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithContext adds a context to the builder
func (app *builder) WithContext(context string) Builder {
	app.context = context
	return app
}

// WithReturn adds a return to the builder
func (app *builder) WithReturn(ret string) Builder {
	app.ret = ret
	return app
}

// Now builds a new Amount instance
func (app *builder) Now() (Amount, error) {
	if app.context == "" {
		return nil, errors.New("the context is mandatory in order to build an Amount instance")
	}

	if app.ret == "" {
		return nil, errors.New("the return is mandatory in order to build an Amount instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.context),
		[]byte(app.ret),
	})

	if err != nil {
		return nil, err
	}

	return createAmount(*pHash, app.context, app.ret), nil
}
