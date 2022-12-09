package instructions

import "errors"

type assignmentBuilder struct {
	variable []byte
	value    Value
}

func createAssignmentBuilder() AssignmentBuilder {
	out := assignmentBuilder{
		variable: nil,
		value:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignmentBuilder) Create() AssignmentBuilder {
	return createAssignmentBuilder()
}

// WithVariable adds a variable to the builder
func (app *assignmentBuilder) WithVariable(variable []byte) AssignmentBuilder {
	app.variable = variable
	return app
}

// WithValue adds a value to the builder
func (app *assignmentBuilder) WithValue(value Value) AssignmentBuilder {
	app.value = value
	return app
}

// Now builds a new Assignment instance
func (app *assignmentBuilder) Now() (Assignment, error) {
	if app.variable == nil {
		return nil, errors.New("the variable is mandatory in order to build an Assignment instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build an Assignment instance")
	}

	return createAssignment(app.variable, app.value), nil
}
