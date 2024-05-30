package modifications

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	insert      string
	delete      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		insert:      "",
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

// WithInsert adds an insert to the builder
func (app *builder) WithInsert(insert string) Builder {
	app.insert = insert
	return app
}

// WithDelete adds a delete to the builder
func (app *builder) WithDelete(delete string) Builder {
	app.delete = delete
	return app
}

// Now builds a new Modification instance
func (app *builder) Now() (Modification, error) {
	data := [][]byte{}
	if app.insert != "" {
		data = append(data, []byte("insert"))
		data = append(data, []byte(app.insert))
	}

	if app.delete != "" {
		data = append(data, []byte("delete"))
		data = append(data, []byte(app.delete))
	}

	if len(data) != 2 {
		return nil, errors.New("the Modification is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.insert),
		[]byte(app.delete),
	})

	if err != nil {
		return nil, err
	}

	if app.insert != "" {
		return createModificationWithInsert(*pHash, app.insert), nil
	}

	return createModificationWithDelete(*pHash, app.delete), nil

}
