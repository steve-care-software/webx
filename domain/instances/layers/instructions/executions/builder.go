package executions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions/merges"
	"github.com/steve-care-software/historydb/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	commit      string
	rollback    string
	cancel      string
	merge       merges.Merge
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		commit:      "",
		rollback:    "",
		cancel:      "",
		merge:       nil,
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
func (app *builder) WithCommit(commit string) Builder {
	app.commit = commit
	return app
}

// WithRollback adds a rollback to the builder
func (app *builder) WithRollback(rollback string) Builder {
	app.rollback = rollback
	return app
}

// WithCancel adds a cancel to the builder
func (app *builder) WithCancel(cancel string) Builder {
	app.cancel = cancel
	return app
}

// WithMerge adds a merge to the builder
func (app *builder) WithMerge(merge merges.Merge) Builder {
	app.merge = merge
	return app
}

// Now builds a new Execution instance
func (app *builder) Now() (Execution, error) {
	bytes := [][]byte{}
	if app.commit != "" {
		bytes = append(bytes, []byte("commit"))
		bytes = append(bytes, []byte(app.commit))
	}

	if app.rollback != "" {
		bytes = append(bytes, []byte("rollback"))
		bytes = append(bytes, []byte(app.rollback))
	}

	if app.cancel != "" {
		bytes = append(bytes, []byte("cancel"))
		bytes = append(bytes, []byte(app.cancel))
	}

	if app.merge != nil {
		bytes = append(bytes, []byte("merge"))
		bytes = append(bytes, app.merge.Hash().Bytes())
	}

	if len(bytes) != 2 {
		return nil, errors.New("the Execution is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(bytes)
	if err != nil {
		return nil, err
	}

	if app.commit != "" {
		return createExecutionWithCommit(*pHash, app.commit), nil
	}

	if app.rollback != "" {
		return createExecutionWithRollback(*pHash, app.rollback), nil
	}

	if app.cancel != "" {
		return createExecutionWithCancel(*pHash, app.cancel), nil
	}

	return createExecutionWithMerge(*pHash, app.merge), nil
}
