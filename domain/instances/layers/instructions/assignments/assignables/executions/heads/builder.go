package heads

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

// WithContext adds the context to the builder
func (app *builder) WithContext(context string) Builder {
	app.context = context
	return app
}

// WithReturn adds a return to the builder
func (app *builder) WithReturn(ret string) Builder {
	app.ret = ret
	return app
}

// Now builds a new Head instance
func (app *builder) Now() (Head, error) {
	if app.context == "" {
		return nil, errors.New("the context is mandatory in order to build an Head instance")
	}

	if app.ret == "" {
		return nil, errors.New("the return is mandatory in order to build an Head instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.context),
		[]byte(app.ret),
	})

	if err != nil {
		return nil, err
	}

	return createHead(*pHash, app.context, app.ret), nil
}
