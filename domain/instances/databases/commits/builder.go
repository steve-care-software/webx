package commits

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions"
)

type builder struct {
	hashAdapter hash.Adapter
	description string
	actions     actions.Actions
	parent      hash.Hash
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		description: "",
		actions:     nil,
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

// WithDescription adds a description to the builder
func (app *builder) WithDescription(description string) Builder {
	app.description = description
	return app
}

// WithActions add an actions to the builder
func (app *builder) WithActions(actions actions.Actions) Builder {
	app.actions = actions
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent hash.Hash) Builder {
	app.parent = parent
	return app
}

// Now builds a new Commit instance
func (app *builder) Now() (Commit, error) {
	if app.description == "" {
		return nil, errors.New("the description is mandatory in order to build a Commit instance")
	}

	if app.actions == nil {
		return nil, errors.New("the actions is mandatory in order to build a Commit instance")
	}

	data := [][]byte{
		[]byte(app.description),
		app.actions.Hash().Bytes(),
	}

	if app.parent != nil {
		data = append(data, app.parent.Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	content := createContent(app.description, app.actions)
	if app.parent != nil {
		return createCommitWithParent(*pHash, content, app.parent), nil
	}

	return createCommit(*pHash, content), nil
}
