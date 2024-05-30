package databases

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits"
	"github.com/steve-care-software/datastencil/domain/instances/databases/heads"
)

type builder struct {
	hashAdapter hash.Adapter
	commit      commits.Commit
	head        heads.Head
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		commit:      nil,
		head:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithCommit adds a commit to the builder
func (app *builder) WithCommit(commit commits.Commit) Builder {
	app.commit = commit
	return app
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head heads.Head) Builder {
	app.head = head
	return app
}

// Now builds a new Database
func (app *builder) Now() (Database, error) {
	if app.commit == nil {
		return nil, errors.New("the commit is mandatory in order to build a Database instance")
	}

	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build a Database instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.commit.Hash().Bytes(),
		app.head.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createDatabase(*pHash, app.commit, app.head), nil
}
