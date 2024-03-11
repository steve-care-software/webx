package reverts

import "github.com/steve-care-software/datastencil/domain/hash"

type builder struct {
	hashAdapter hash.Adapter
	index       string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		index:       "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index string) Builder {
	app.index = index
	return app
}

// Now builds a new Revert instance
func (app *builder) Now() (Revert, error) {
	data := [][]byte{
		[]byte("revert"),
	}

	if app.index != "" {
		data = append(data, []byte(app.index))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.index != "" {
		return createRevertWithIndex(*pHash, app.index), nil
	}

	return createRevert(*pHash), nil
}
