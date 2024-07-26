package lists

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/lists/fetches"
)

type builder struct {
	hashAdapter hash.Adapter
	fetch       fetches.Fetch
	length      string
	create      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		fetch:       nil,
		length:      "",
		create:      "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithFetch adds a fetch to the builder
func (app *builder) WithFetch(fetch fetches.Fetch) Builder {
	app.fetch = fetch
	return app
}

// WithLength adds a length to the builder
func (app *builder) WithLength(length string) Builder {
	app.length = length
	return app
}

// WithCreate adds a create to the builder
func (app *builder) WithCreate(create string) Builder {
	app.create = create
	return app
}

// Now builds a new List instance
func (app *builder) Now() (List, error) {
	data := [][]byte{}
	if app.fetch != nil {
		data = append(data, []byte("fetch"))
		data = append(data, app.fetch.Hash().Bytes())
	}

	if app.length != "" {
		data = append(data, []byte("length"))
		data = append(data, []byte(app.length))
	}

	if app.create != "" {
		data = append(data, []byte("create"))
		data = append(data, []byte(app.create))
	}

	if len(data) != 2 {
		return nil, errors.New("the List is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.fetch != nil {
		return createListWithFetch(*pHash, app.fetch), nil
	}

	if app.length != "" {
		return createListWithLength(*pHash, app.length), nil
	}

	return createListWithCreate(*pHash, app.create), nil
}
