package files

import (
	"errors"

	"github.com/steve-care-software/webx/engine/cursors/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	close       string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		close:       "",
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

// Now builds a new File instance
func (app *builder) Now() (File, error) {
	data := [][]byte{}
	if app.close != "" {
		data = append(data, []byte("close"))
		data = append(data, []byte(app.close))
	}

	if len(data) != 2 {
		return nil, errors.New("the File is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	return createFileWithClose(*pHash, app.close), nil
}
