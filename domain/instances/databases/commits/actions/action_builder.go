package actions

import (
	"errors"
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/values"
)

type actionBuilder struct {
	hashAdapter hash.Adapter
	path        []string
	insert      values.Value
	isDelete    bool
}

func createActionBuilder(
	hashAdapter hash.Adapter,
) ActionBuilder {
	out := actionBuilder{
		hashAdapter: hashAdapter,
		path:        nil,
		insert:      nil,
		isDelete:    false,
	}

	return &out
}

// Create initializes the builder
func (app *actionBuilder) Create() ActionBuilder {
	return createActionBuilder(
		app.hashAdapter,
	)
}

// WithPath adds a path to the builder
func (app *actionBuilder) WithPath(path []string) ActionBuilder {
	app.path = path
	return app
}

// WithInsert adds an insert to the builder
func (app *actionBuilder) WithInsert(insert values.Value) ActionBuilder {
	app.insert = insert
	return app
}

// IsDelete flags the builder as a delete
func (app *actionBuilder) IsDelete() ActionBuilder {
	app.isDelete = true
	return app
}

// Now builds a new Action instance
func (app *actionBuilder) Now() (Action, error) {
	data := [][]byte{}
	if app.insert != nil {
		data = append(data, []byte("insert"))
		data = append(data, app.insert.Hash().Bytes())
	}

	if app.isDelete {
		data = append(data, []byte("delete"))
	}

	amount := len(data)
	if amount != 1 && amount != 2 {
		return nil, errors.New("the Action is invalid")
	}

	path := filepath.Join(app.path...)
	data = append(data, []byte(path))
	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.insert != nil {
		content := createContentWithInsert(app.insert)
		return createAction(*pHash, app.path, content), nil
	}

	content := createContentWithDelete()
	return createAction(*pHash, app.path, content), nil
}
