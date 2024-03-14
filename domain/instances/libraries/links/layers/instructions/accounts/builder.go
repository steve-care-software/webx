package accounts

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/accounts/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/libraries/links/layers/instructions/accounts/updates"
)

type builder struct {
	hashAdapter hash.Adapter
	insert      inserts.Insert
	update      updates.Update
	delete      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		insert:      nil,
		update:      nil,
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

// WithInsert adds an insert to the builder
func (app *builder) WithInsert(insert inserts.Insert) Builder {
	app.insert = insert
	return app
}

// WithUpdate adds an update to the builder
func (app *builder) WithUpdate(update updates.Update) Builder {
	app.update = update
	return app
}

// WithDelete adds a delete to the builder
func (app *builder) WithDelete(delete string) Builder {
	app.delete = delete
	return app
}

// Now builds a new Account instance
func (app *builder) Now() (Account, error) {
	data := [][]byte{}
	if app.insert != nil {
		data = append(data, []byte("insert"))
		data = append(data, app.insert.Hash().Bytes())
	}

	if app.update != nil {
		data = append(data, []byte("update"))
		data = append(data, app.update.Hash().Bytes())
	}

	if app.delete != "" {
		data = append(data, []byte("delete"))
		data = append(data, []byte(app.delete))
	}

	if len(data) != 2 {
		return nil, errors.New("the Account is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.insert != nil {
		return createAccountWithInsert(*pHash, app.insert), nil
	}

	if app.update != nil {
		return createAccountWithUpdate(*pHash, app.update), nil
	}

	return createAccountWithDelete(*pHash, app.delete), nil
}
