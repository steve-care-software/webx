package actions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	path        string
	insert      string
	isDelete    bool
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		path:        "",
		insert:      "",
		isDelete:    false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path string) Builder {
	app.path = path
	return app
}

// WithInsert adds an insert to the builder
func (app *builder) WithInsert(insert string) Builder {
	app.insert = insert
	return app
}

// IsDelete flags the builder as delete
func (app *builder) IsDelete() Builder {
	app.isDelete = true
	return app
}

// Now builds a new Action instance
func (app *builder) Now() (Action, error) {
	if app.path == "" {
		return nil, errors.New("the path is mandatory in order to build an Action instance")
	}

	data := [][]byte{}
	if app.insert != "" {
		data = append(data, []byte("insert"))
		data = append(data, []byte(app.insert))
	}

	if app.isDelete {
		data = append(data, []byte("isDelete"))
	}

	amount := len(data)
	if amount != 1 && amount != 2 {
		return nil, errors.New("the Action is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.insert != "" {
		return createActionWithInsert(*pHash, app.path, app.insert), nil
	}

	return createActionWithDelete(*pHash, app.path), nil
}
