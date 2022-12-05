package programs

import "errors"

type assignmentBuilder struct {
	pIndex *uint
	name   []byte
	value  Value
}

func createAssignmentBuilder() AssignmentBuilder {
	out := assignmentBuilder{
		pIndex: nil,
		name:   nil,
		value:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignmentBuilder) Create() AssignmentBuilder {
	return createAssignmentBuilder()
}

// WithIndex adds an index to the builder
func (app *assignmentBuilder) WithIndex(index uint) AssignmentBuilder {
	app.pIndex = &index
	return app
}

// WithName adds a name to the builder
func (app *assignmentBuilder) WithName(name []byte) AssignmentBuilder {
	app.name = name
	return app
}

// WithValue adds a value to the builder
func (app *assignmentBuilder) WithValue(value Value) AssignmentBuilder {
	app.value = value
	return app
}

// Now builds a new Assignment instance
func (app *assignmentBuilder) Now() (Assignment, error) {
	if app.name == nil {
		return nil, errors.New("the name is mandatory in order to build an Assignment instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build an Assignment instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build an Assignment instance")
	}

	return createAssignment(*app.pIndex, app.name, app.value), nil
}
