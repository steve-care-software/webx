package grammars

import (
	"errors"

	"github.com/steve-care-software/webx/databases/domain/cryptography/hash"
)

type everythingBuilder struct {
	hashAdapter hash.Adapter
	name        string
	exception   Token
	escape      Token
}

func createEverythingBuilder(
	hashAdapter hash.Adapter,
) EverythingBuilder {
	out := everythingBuilder{
		hashAdapter: hashAdapter,
		name:        "",
		exception:   nil,
		escape:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *everythingBuilder) Create() EverythingBuilder {
	return createEverythingBuilder(
		app.hashAdapter,
	)
}

// WithName adds a name to the builder
func (app *everythingBuilder) WithName(name string) EverythingBuilder {
	app.name = name
	return app
}

// WithException adds an exception to the builder
func (app *everythingBuilder) WithException(exception Token) EverythingBuilder {
	app.exception = exception
	return app
}

// WithEscape adds an escape to the builder
func (app *everythingBuilder) WithEscape(escape Token) EverythingBuilder {
	app.escape = escape
	return app
}

// Now builds a new Everything instance
func (app *everythingBuilder) Now() (Everything, error) {
	if app.exception == nil {
		return nil, errors.New("the exception is mandatory in order to build an Everything instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Everything instance")
	}

	data := [][]byte{
		[]byte(app.name),
		app.exception.Hash().Bytes(),
	}

	if app.escape != nil {
		data = append(data, app.escape.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.escape != nil {
		return createEverythingWithEscape(*pHash, app.name, app.exception, app.escape), nil
	}

	return createEverything(*pHash, app.name, app.exception), nil
}
