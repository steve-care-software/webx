package states

import (
	"errors"

	"github.com/steve-care-software/webx/engine/databases/bytes/domain/headers/states/containers"
)

type stateBuilder struct {
	isDeleted  bool
	containers containers.Containers
}

func createStateBuilder() StateBuilder {
	out := stateBuilder{
		isDeleted:  false,
		containers: nil,
	}

	return &out
}

// Create initializes the builder
func (app *stateBuilder) Create() StateBuilder {
	return createStateBuilder()
}

// WithContainers add containers to the builder
func (app *stateBuilder) WithContainers(containers containers.Containers) StateBuilder {
	app.containers = containers
	return app
}

// IsDeleted flags the builder as deleted
func (app *stateBuilder) IsDeleted() StateBuilder {
	app.isDeleted = true
	return app
}

// Now builds a new State instance
func (app *stateBuilder) Now() (State, error) {
	if app.containers == nil {
		return nil, errors.New("the containers is mandatory in order to build a State instance")
	}

	if app.isDeleted {
		return createStateWithDeleted(app.containers), nil
	}

	return createState(app.containers), nil
}
