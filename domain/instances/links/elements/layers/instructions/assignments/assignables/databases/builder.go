package databases

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/actions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/databases"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/modifications"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/retrieves"
)

type builder struct {
	hashAdapter  hash.Adapter
	db           databases.Database
	commit       commits.Commit
	action       actions.Action
	modification modifications.Modification
	delete       deletes.Delete
	retrieve     retrieves.Retrieve
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:  hashAdapter,
		db:           nil,
		commit:       nil,
		action:       nil,
		modification: nil,
		delete:       nil,
		retrieve:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithDatabase adds a database to the builder
func (app *builder) WithDatabase(database databases.Database) Builder {
	app.db = database
	return app
}

// WithCommit adds a commit to the builder
func (app *builder) WithCommit(commit commits.Commit) Builder {
	app.commit = commit
	return app
}

// WithAction adds an action to the builder
func (app *builder) WithAction(action actions.Action) Builder {
	app.action = action
	return app
}

// WithModification adds a modification to the builder
func (app *builder) WithModification(modification modifications.Modification) Builder {
	app.modification = modification
	return app
}

// WithDelete adds a delete to the builder
func (app *builder) WithDelete(delete deletes.Delete) Builder {
	app.delete = delete
	return app
}

// WithRetrieve adds a retrieve to the builder
func (app *builder) WithRetrieve(retrieve retrieves.Retrieve) Builder {
	app.retrieve = retrieve
	return app
}

// Now builds a new Database instance
func (app *builder) Now() (Database, error) {
	data := [][]byte{}
	if app.db != nil {
		data = append(data, []byte("database"))
		data = append(data, app.db.Hash().Bytes())
	}

	if app.commit != nil {
		data = append(data, []byte("commit"))
		data = append(data, app.commit.Hash().Bytes())
	}

	if app.action != nil {
		data = append(data, []byte("action"))
		data = append(data, app.action.Hash().Bytes())
	}

	if app.modification != nil {
		data = append(data, []byte("modification"))
		data = append(data, app.modification.Hash().Bytes())
	}

	if app.delete != nil {
		data = append(data, []byte("delete"))
		data = append(data, app.delete.Hash().Bytes())
	}

	if app.retrieve != nil {
		data = append(data, []byte("retrieve"))
		data = append(data, app.retrieve.Hash().Bytes())
	}

	if len(data) != 2 {
		return nil, errors.New("the Database is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.db != nil {
		return createDatabaseWithDatabase(*pHash, app.db), nil
	}

	if app.commit != nil {
		return createDatabaseWithCommit(*pHash, app.commit), nil
	}

	if app.action != nil {
		return createDatabaseWithAction(*pHash, app.action), nil
	}

	if app.modification != nil {
		return createDatabaseWithModification(*pHash, app.modification), nil
	}

	if app.delete != nil {
		return createDatabaseWithDelete(*pHash, app.delete), nil
	}

	return createDatabaseWithRetrieve(*pHash, app.retrieve), nil
}
