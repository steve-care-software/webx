package databases

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	save        string
	delete      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		save:        "",
		delete:      "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithSave adds a save to the builder
func (app *builder) WithSave(save string) Builder {
	app.save = save
	return app
}

// WithDelete adds a delete to the builder
func (app *builder) WithDelete(delete string) Builder {
	app.delete = delete
	return app
}

// Now builds a new Database instance
func (app *builder) Now() (Database, error) {
	data := [][]byte{}
	if app.save != "" {
		data = append(data, []byte("save"))
		data = append(data, []byte(app.save))
	}

	if app.delete != "" {
		data = append(data, []byte("delete"))
		data = append(data, []byte(app.delete))
	}

	if len(data) != 2 {
		return nil, errors.New("the Database is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.save != "" {
		return createDatabaseWithSave(*pHash, app.save), nil
	}

	return createDatabaseWithDelete(*pHash, app.delete), nil
}
