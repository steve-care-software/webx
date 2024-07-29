package listers

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/retrievals"
)

type builder struct {
	keyname   string
	retrieval retrievals.Retrieval
}

func createBuilder() Builder {
	out := builder{
		keyname:   "",
		retrieval: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithKeyname adds a keyname to the builder
func (app *builder) WithKeyname(keyname string) Builder {
	app.keyname = keyname
	return app
}

// WithRetrieval adds a retrieval to the builder
func (app *builder) WithRetrieval(retrieval retrievals.Retrieval) Builder {
	app.retrieval = retrieval
	return app
}

// Now builds a new Lister instance
func (app *builder) Now() (Lister, error) {
	if app.keyname == "" {
		return nil, errors.New("the keyname is mandatory in order to build a Lister instance")
	}

	if app.retrieval == nil {
		return nil, errors.New("the retrieval is mandatory in order to build a Lister instance")
	}

	return createLister(
		app.keyname,
		app.retrieval,
	), nil
}
