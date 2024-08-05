package states

import (
	"errors"

	"github.com/steve-care-software/webx/engine/bytes/domain/pointers"
)

type stateBuilder struct {
	message   string
	isDeleted bool
	pointers  pointers.Pointers
}

func createStateBuilder() StateBuilder {
	out := stateBuilder{
		message:   "",
		isDeleted: false,
		pointers:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *stateBuilder) Create() StateBuilder {
	return createStateBuilder()
}

// WithMessage adds a message to the builder
func (app *stateBuilder) WithMessage(message string) StateBuilder {
	app.message = message
	return app
}

// WithPointers adds pointers to the builder
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
	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build a State instance")
	}

	if app.pointers != nil {
		return createStateWithPointers(app.message, app.isDeleted, app.pointers), nil
	}

	return createState(app.message, app.isDeleted), nil
}
