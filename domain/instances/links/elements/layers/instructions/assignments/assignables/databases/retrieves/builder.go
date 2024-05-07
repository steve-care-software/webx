package retrieves

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	exists      string
	retrieve    string
	isList      bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		exists:      "",
		retrieve:    "",
		isList:      false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithExists adds an exists to the builder
func (app *builder) WithExists(exists string) Builder {
	app.exists = exists
	return app
}

// WithRetrieve adds a retrieve to the builder
func (app *builder) WithRetrieve(retrieve string) Builder {
	app.retrieve = retrieve
	return app
}

// IsList flags the builder as a list
func (app *builder) IsList() Builder {
	app.isList = true
	return app
}

// Now builds a new Retrieve instance
func (app *builder) Now() (Retrieve, error) {
	data := [][]byte{}
	if app.exists != "" {
		data = append(data, []byte("exists"))
		data = append(data, []byte(app.exists))
	}

	if app.retrieve != "" {
		data = append(data, []byte("retrieve"))
		data = append(data, []byte(app.retrieve))
	}

	if app.isList {
		data = append(data, []byte("isList"))
	}

	amount := len(data)
	if amount != 1 && amount != 2 {
		return nil, errors.New("the Retrieve is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.exists != "" {
		return createRetrieveWithExists(*pHash, app.exists), nil
	}

	if app.retrieve != "" {
		return createRetrieveWithRetrieve(*pHash, app.retrieve), nil
	}

	return createRetrieveWithList(*pHash), nil
}
