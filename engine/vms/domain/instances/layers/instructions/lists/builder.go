package lists

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/lists/deletes"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/lists/inserts"
)

type builder struct {
	hashAdapter hash.Adapter
	insert      inserts.Insert
	delete      deletes.Delete
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		insert:      nil,
		delete:      nil,
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
func (app *builder) WithInsert(insert inserts.Insert) Builder {
	app.insert = insert
	return app
}

// WithDelete adds a delete to the builder
func (app *builder) WithDelete(delete deletes.Delete) Builder {
	app.delete = delete
	return app
}

// Now builds a new List instance
func (app *builder) Now() (List, error) {
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
		return nil, errors.New("the List is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.insert != nil {
		return createListWithInsert(*pHash, app.insert), nil
	}

	return createListWithDelete(*pHash, app.delete), nil
}
