package contents

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/actions"
	"github.com/steve-care-software/datastencil/domain/instances/commits/contents/previous"
)

type builder struct {
	hashAdapter hash.Adapter
	action      actions.Action
	previous    previous.Previous
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		action:      nil,
		previous:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithAction adds an action to the builder
func (app *builder) WithAction(action actions.Action) Builder {
	app.action = action
	return app
}

// WithPrevious adds a previous to the builder
func (app *builder) WithPrevious(previous previous.Previous) Builder {
	app.previous = previous
	return app
}

// Now builds a new Content instance
func (app *builder) Now() (Content, error) {
	if app.action == nil {
		return nil, errors.New("the action is mandatory in order to build a Content instance")
	}

	if app.previous == nil {
		return nil, errors.New("the previous is mandatory in order to build a Content instance")
	}

	pHash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.action.Hash().Bytes(),
		app.previous.Hash().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createContent(*pHash, app.action, app.previous), nil
}
