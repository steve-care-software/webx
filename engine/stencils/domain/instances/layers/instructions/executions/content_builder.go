package executions

import (
	"errors"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/instructions/executions/merges"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	commit      string
	rollback    string
	cancel      string
	merge       merges.Merge
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter: hashAdapter,
		commit:      "",
		rollback:    "",
		cancel:      "",
		merge:       nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(
		app.hashAdapter,
	)
}

// WithCommit adds a commit to the builder
func (app *contentBuilder) WithCommit(commit string) ContentBuilder {
	app.commit = commit
	return app
}

// WithRollback adds a rollback to the builder
func (app *contentBuilder) WithRollback(rollback string) ContentBuilder {
	app.rollback = rollback
	return app
}

// WithCancel adds a cancel to the builder
func (app *contentBuilder) WithCancel(cancel string) ContentBuilder {
	app.cancel = cancel
	return app
}

// WithMerge adds a merge to the builder
func (app *contentBuilder) WithMerge(merge merges.Merge) ContentBuilder {
	app.merge = merge
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
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
		return nil, errors.New("the Content is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(bytes)
	if err != nil {
		return nil, err
	}

	if app.commit != "" {
		return createContentWithCommit(*pHash, app.commit), nil
	}

	if app.rollback != "" {
		return createContentWithRollback(*pHash, app.rollback), nil
	}

	if app.cancel != "" {
		return createContentWithCancel(*pHash, app.cancel), nil
	}

	return createContentWithMerge(*pHash, app.merge), nil
}
