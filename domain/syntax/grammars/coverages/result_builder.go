package coverages

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/trees"
)

type resultBuilder struct {
	tree  trees.Tree
	error string
}

func createResultBuilder() ResultBuilder {
	out := resultBuilder{
		tree:  nil,
		error: "",
	}

	return &out
}

// Create initializes the builder
func (app *resultBuilder) Create() ResultBuilder {
	return createResultBuilder()
}

// WithTree adds a tree to the builder
func (app *resultBuilder) WithTree(tree trees.Tree) ResultBuilder {
	app.tree = tree
	return app
}

// WithError adds an error to the builder
func (app *resultBuilder) WithError(error string) ResultBuilder {
	app.error = error
	return app
}

// Now builds a new Result instance
func (app *resultBuilder) Now() (Result, error) {
	if app.tree != nil {
		return createResultWithTree(app.tree), nil
	}

	if app.error != "" {
		return createResultWithError(app.error), nil
	}

	return nil, errors.New("the Result is invalid")
}
