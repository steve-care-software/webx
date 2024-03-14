package databases

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/inserts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/databases/reverts"
)

type builder struct {
	hashAdapter hash.Adapter
	insert      inserts.Insert
	delete      deletes.Delete
	commit      string
	cancel      string
	revert      reverts.Revert
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		insert:      nil,
		delete:      nil,
		commit:      "",
		cancel:      "",
		revert:      nil,
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

// WithCommit adds a commit to the builder
func (app *builder) WithCommit(commit string) Builder {
	app.commit = commit
	return app
}

// WithCancel adds a cancel to the builder
func (app *builder) WithCancel(cancel string) Builder {
	app.cancel = cancel
	return app
}

// WithRevert adds a revert to the builder
func (app *builder) WithRevert(revert reverts.Revert) Builder {
	app.revert = revert
	return app
}

// Now builds a new Database instance
func (app *builder) Now() (Database, error) {
	data := [][]byte{}
	if app.insert != nil {
		data = append(data, []byte("insert"))
		data = append(data, app.insert.Hash().Bytes())
	}

	if app.delete != nil {
		data = append(data, []byte("delete"))
		data = append(data, app.delete.Hash().Bytes())
	}

	if app.commit != "" {
		data = append(data, []byte("commit"))
		data = append(data, []byte(app.commit))
	}

	if app.cancel != "" {
		data = append(data, []byte("cancel"))
		data = append(data, []byte(app.cancel))
	}

	if app.revert != nil {
		data = append(data, []byte("revert"))
		data = append(data, app.revert.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.insert != nil {
		return createDatabaseWithInsert(*pHash, app.insert), nil
	}

	if app.delete != nil {
		return createDatabaseWithDelete(*pHash, app.delete), nil
	}

	if app.commit != "" {
		return createDatabaseWithCommit(*pHash, app.commit), nil
	}

	if app.cancel != "" {
		return createDatabaseWithCancel(*pHash, app.cancel), nil
	}

	return createDatabaseWithRevert(*pHash, app.revert), nil
}
