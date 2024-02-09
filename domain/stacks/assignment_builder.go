package stacks

import "errors"

type assignmentBuilder struct {
	name       string
	assignable Assignable
}

func createAssignmentBuilder() AssignmentBuilder {
	out := assignmentBuilder{
		name:       "",
		assignable: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignmentBuilder) Create() AssignmentBuilder {
	return createAssignmentBuilder()
}

// WithName adds a name to the builder
func (app *assignmentBuilder) WithName(name string) AssignmentBuilder {
	app.name = name
	return app
}

// WithAssignable adds an assignable to the builder
func (app *assignmentBuilder) WithAssignable(assignable Assignable) AssignmentBuilder {
	app.assignable = assignable
	return app
}

// Now builds a new Assignment instance
func (app *assignmentBuilder) Now() (Assignment, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build an Assignment instance")
	}

	if app.assignable == nil {
		return nil, errors.New("the assignable is mandatory in order to build an Assignment instance")
	}

	return createAssignment(app.name, app.assignable), nil
}
