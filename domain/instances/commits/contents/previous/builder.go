package previous

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/actions"
)

type builder struct {
	hashAdapter hash.Adapter
	root        actions.Actions
	prev        Previous
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		root:        nil,
		prev:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root actions.Actions) Builder {
	app.root = root
	return app
}

// WithPrevious adds a previous to the builder
func (app *builder) WithPrevious(previous Previous) Builder {
	app.prev = previous
	return app
}

// Now builds a new Previous instance
func (app *builder) Now() (Previous, error) {
	data := [][]byte{}
	if app.root != nil {
		data = append(data, []byte("root"))
		data = append(data, app.root.Hash().Bytes())
	}

	if app.prev != nil {
		data = append(data, []byte("previous"))
		data = append(data, app.prev.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Previous is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.root != nil {
		return createPreviousWithRoot(*pHash, app.root), nil
	}

	return createPreviousWithPrevious(*pHash, app.prev), nil
}
