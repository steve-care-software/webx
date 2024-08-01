package states

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/pointers"
)

type stateBuilder struct {
	isDeleted bool
	pointers  pointers.Pointers
}

func createStateBuilder() StateBuilder {
	out := stateBuilder{
		isDeleted: false,
		pointers:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *stateBuilder) Create() StateBuilder {
	return createStateBuilder()
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
	if app.isDeleted && app.pointers != nil {
		return createStateWithPointersAndDeleted(app.pointers), nil
	}

	if app.pointers != nil {
		return createStateWithPointers(app.pointers), nil
	}

	if app.isDeleted {
		return createStateWithDeleted(app.pointers), nil
	}

	return createState(), nil
}
