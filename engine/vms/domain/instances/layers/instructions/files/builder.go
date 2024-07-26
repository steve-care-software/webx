package files

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/entities/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	close       string
	delete      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		close:       "",
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

// WithClose adds a close to the builder
func (app *builder) WithClose(close string) Builder {
	app.close = close
	return app
}

// WithDelete adds a delete to the builder
func (app *builder) WithDelete(del string) Builder {
	app.delete = del
	return app
}

// Now builds a new File instance
func (app *builder) Now() (File, error) {
	data := [][]byte{}
	if app.close != "" {
		data = append(data, []byte("close"))
		data = append(data, []byte(app.close))
	}

	if app.delete != "" {
		data = append(data, []byte("delete"))
		data = append(data, []byte(app.delete))
	}

	if len(data) != 2 {
		return nil, errors.New("the File is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.close != "" {
		return createFileWithClose(*pHash, app.close), nil
	}

	return createFileWithDelete(*pHash, app.delete), nil
}
