package commits

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	actions     actions.Actions
	previous    Commit
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter: hashAdapter,
		actions:     nil,
		previous:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(
		app.hashAdapter,
	)
}

// WithActions adds an actions to the builder
func (app *contentBuilder) WithActions(actions actions.Actions) ContentBuilder {
	app.actions = actions
	return app
}

// WithPrevious adds a previous to the builder
func (app *contentBuilder) WithPrevious(previous Commit) ContentBuilder {
	app.previous = previous
	return app
}

// Now builds a new Content instance
func (app *contentBuilder) Now() (Content, error) {
	if app.actions == nil {
		return nil, errors.New("the actions is mandatory in order to build a Content instance")
	}

	data := [][]byte{
		app.actions.Hash().Bytes(),
	}

	if app.previous != nil {
		data = append(data, app.previous.Hash().Bytes())
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.previous != nil {
		return createContentWithPrevious(*pHash, app.actions, app.previous), nil
	}

	return createContent(*pHash, app.actions), nil
}
