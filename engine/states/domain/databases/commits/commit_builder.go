package commits

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/databases/commits/executions"
	"github.com/steve-care-software/webx/engine/states/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	executions  executions.Executions
	parent      *hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	return &builder{
		hashAdapter: hashAdapter,
		executions:  nil,
		parent:      nil,
	}
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithExecutions adds an executions to the builder
func (app *builder) WithExecutions(executions executions.Executions) Builder {
	app.executions = executions
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent hash.Hash) Builder {
	app.parent = &parent
	return app
}

// Now builds a new Commit instance
func (app *builder) Now() (Commit, error) {
	if app.executions == nil {
		return nil, errors.New("the executions are mandatory in order to build a Commit instance")
	}

	data := [][]byte{
		app.executions.Hash().Bytes(),
	}

	if app.parent != nil {
		data = append(data, app.parent.Bytes())
	}

	commitHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		return createCommitWithParent(*commitHash, app.executions, *app.parent), nil
	}

	return createCommit(*commitHash, app.executions), nil
}
