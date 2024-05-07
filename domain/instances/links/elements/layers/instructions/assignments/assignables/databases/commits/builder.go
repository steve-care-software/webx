package commits

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	description string
	actions     string
	parent      string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		description: "",
		actions:     "",
		parent:      "",
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

// WithActions adds an actions to the builder
func (app *builder) WithActions(actions string) Builder {
	app.actions = actions
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent string) Builder {
	app.parent = parent
	return app
}

// Now builds a new Commit instance
func (app *builder) Now() (Commit, error) {
	if app.description == "" {
		return nil, errors.New("the description is mandatory in order to build a Commit instance")
	}

	if app.actions == "" {
		return nil, errors.New("the actions is mandatory in order to build a Commit instance")
	}

	data := [][]byte{
		[]byte(app.description),
		[]byte(app.actions),
	}

	if app.parent != "" {
		data = append(data, []byte(app.parent))
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.parent != "" {
		return createCommitWithParent(*pHash, app.description, app.actions, app.parent), nil
	}

	return createCommit(*pHash, app.description, app.actions), nil
}
