package commits

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/queries"
)

type builder struct {
	hashAdapter hash.Adapter
	queries     queries.Queries
	parent      hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		queries:     nil,
		parent:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithQueries add queries to the builder
func (app *builder) WithQueries(queries queries.Queries) Builder {
	app.queries = queries
	return app
}

// WithParent add parent to the builder
func (app *builder) WithParent(parent hash.Hash) Builder {
	app.parent = parent
	return app
}

// Now builds a new Commit instance
func (app *builder) Now() (Commit, error) {
	if app.queries == nil {
		return nil, errors.New("the queries is mandatory in order to build a Commit instance")
	}

	data := [][]byte{
		app.queries.Hash().Bytes(),
	}

	if app.parent != nil {
		data = append(data, app.parent.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		return createCommitWithParent(*pHash, app.queries, app.parent), nil
	}

	return createCommit(*pHash, app.queries), nil
}
