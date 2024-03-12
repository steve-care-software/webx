package actions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/actions/pointers"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/actions/resources"
)

type actionBuilder struct {
	hashAdapter hash.Adapter
	insert      resources.Resource
	delete      pointers.Pointer
}

func createActionBuilder(
	hashAdapter hash.Adapter,
) ActionBuilder {
	out := actionBuilder{
		hashAdapter: hashAdapter,
		insert:      nil,
		delete:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *actionBuilder) Create() ActionBuilder {
	return createActionBuilder(
		app.hashAdapter,
	)
}

// WithInsert adds an insert to the builder
func (app *actionBuilder) WithInsert(insert resources.Resource) ActionBuilder {
	app.insert = insert
	return app
}

// WithDelete adds a delete to the builder
func (app *actionBuilder) WithDelete(delete pointers.Pointer) ActionBuilder {
	app.delete = delete
	return app
}

// Now builds a new Action instance
func (app *actionBuilder) Now() (Action, error) {
	data := [][]byte{}
	if app.insert != nil {
		data = append(data, []byte("insert"))
		data = append(data, app.insert.Hash().Bytes())
	}

	if app.delete != nil {
		data = append(data, []byte("delete"))
		data = append(data, app.delete.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Action is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.insert != nil {
		return createActionWithInsert(*pHash, app.insert), nil
	}

	return createActionWithDelete(*pHash, app.delete), nil
}
