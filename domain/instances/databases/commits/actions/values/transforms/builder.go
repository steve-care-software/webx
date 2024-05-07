package transforms

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	query       []byte
	bytes       []byte
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		query:       nil,
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

// WithQuery adds a query to the builder
func (app *builder) WithQuery(query []byte) Builder {
	app.query = query
	return app
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes []byte) Builder {
	app.bytes = bytes
	return app
}

// Now builds a new Transform instance
func (app *builder) Now() (Transform, error) {
	if app.query != nil && len(app.query) <= 0 {
		app.query = nil
	}

	if app.bytes != nil && len(app.bytes) <= 0 {
		app.bytes = nil
	}

	if app.query == nil {
		return nil, errors.New("the query is mandatory in order to build a Transform instance")
	}

	if app.bytes == nil {
		return nil, errors.New("the bytes is mandatory in order to build a Transform instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.query,
		app.bytes,
	})

	if err != nil {
		return nil, err
	}

	return createTransform(*pHash, app.query, app.bytes), nil
}
