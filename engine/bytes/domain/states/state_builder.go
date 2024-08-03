package states

import (
	"github.com/steve-care-software/webx/engine/bytes/domain/states/pointers"
	"github.com/steve-care-software/webx/engine/bytes/domain/states/pointers/delimiters"
)

type stateBuilder struct {
	isDeleted bool
	root      delimiters.Delimiter
	pointers  pointers.Pointers
}

func createStateBuilder() StateBuilder {
	out := stateBuilder{
		isDeleted: false,
		root:      nil,
		pointers:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *stateBuilder) Create() StateBuilder {
	return createStateBuilder()
}

// WithRoot add root to the builder
func (app *stateBuilder) WithRoot(root delimiters.Delimiter) StateBuilder {
	app.root = root
	return app
}

// WithPointers add pointers to the builder
func (app *stateBuilder) WithPointers(pointers pointers.Pointers) StateBuilder {
	app.pointers = pointers
	return app
}

// IsDeleted flags the builder as deleted
func (app *stateBuilder) IsDeleted() StateBuilder {
	app.isDeleted = true
	return app
}

// Now builds a new State instance
func (app *stateBuilder) Now() (State, error) {
	if app.root != nil && app.pointers != nil {
		return createStateWithRootAndPointers(app.isDeleted, app.root, app.pointers), nil
	}

	if app.root != nil {
		return createStateWithRoot(app.isDeleted, app.root), nil
	}

	if app.pointers != nil {
		return createStateWithPointers(app.isDeleted, app.pointers), nil
	}

	return createState(app.isDeleted), nil
}
