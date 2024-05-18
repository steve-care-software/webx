package stacks

import "errors"

type assignablesBuilder struct {
	list []Assignable
}

func createAssignablesBuilder() AssignablesBuilder {
	out := assignablesBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignablesBuilder) Create() AssignablesBuilder {
	return createAssignablesBuilder()
}

// WithList adds a list to the builder
func (app *assignablesBuilder) WithList(list []Assignable) AssignablesBuilder {
	app.list = list
	return app
}

// Now builds a new Assignables instance
func (app *assignablesBuilder) Now() (Assignables, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Assignable in order to build an Assignables instance")
	}

	return createAssignables(app.list), nil
}
