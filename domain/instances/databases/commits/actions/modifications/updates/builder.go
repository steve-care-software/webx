package updates

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
)

type builder struct {
	hashAdapter hash.Adapter
	delete      deletes.Delete
	bytes       []byte
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		delete:      nil,
		bytes:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithDelete adds a delete to the builder
func (app *builder) WithDelete(del deletes.Delete) Builder {
	app.delete = del
	return app
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes []byte) Builder {
	app.bytes = bytes
	return app
}

// Now builds a new Update instance
func (app *builder) Now() (Update, error) {
	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.bytes == nil {
		return nil, errors.New("the bytes is mandatory in order to build an Updare instance")
	}

	if app.delete == nil {
		return nil, errors.New("the delete is mandatory in order to build an Updare instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.bytes,
		app.delete.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createUpdate(*pHash, app.delete, app.bytes), nil
}
