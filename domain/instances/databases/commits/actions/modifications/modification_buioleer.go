package modifications

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/updates"
)

type modificationBuilder struct {
	hashAdapter hash.Adapter
	insert      []byte
	update      updates.Update
	delete      deletes.Delete
}

func createModificationBuilder(
	hashAdapter hash.Adapter,
) ModificationBuilder {
	out := modificationBuilder{
		hashAdapter: hashAdapter,
		insert:      nil,
		update:      nil,
		delete:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *modificationBuilder) Create() ModificationBuilder {
	return createModificationBuilder(
		app.hashAdapter,
	)
}

// WithInsert adds an insert to the builder
func (app *modificationBuilder) WithInsert(insert []byte) ModificationBuilder {
	app.insert = insert
	return app
}

// WithUpdate adds an update to the builder
func (app *modificationBuilder) WithUpdate(update updates.Update) ModificationBuilder {
	app.update = update
	return app
}

// WithDelete adds a delete to the builder
func (app *modificationBuilder) WithDelete(delete deletes.Delete) ModificationBuilder {
	app.delete = delete
	return app
}

// Now builds a new Modification instance
func (app *modificationBuilder) Now() (Modification, error) {
	data := [][]byte{}
	if app.insert != nil {
		data = append(data, []byte("insert"))
		data = append(data, []byte(app.insert))
	}

	if app.update != nil {
		data = append(data, []byte("update"))
		data = append(data, []byte(app.update.Hash().Bytes()))
	}

	if app.delete != nil {
		data = append(data, []byte("delete"))
		data = append(data, []byte(app.delete.Hash().Bytes()))
	}

	if len(data) != 2 {
		return nil, errors.New("the Modification is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.insert != nil {
		return createModificationWithInsert(*pHash, app.insert), nil
	}

	if app.update != nil {
		return createModificationWithUpdate(*pHash, app.update), nil
	}

	return createModificationWithDelete(*pHash, app.delete), nil
}
