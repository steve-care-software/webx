package stacks

import (
	"errors"
	"fmt"
)

type assignmentsBuilder struct {
	list []Assignment
}

func createAssignmentsBuilder() AssignmentsBuilder {
	out := assignmentsBuilder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *assignmentsBuilder) Create() AssignmentsBuilder {
	return createAssignmentsBuilder()
}

// WithList adds a list to the builder
func (app *assignmentsBuilder) WithList(list []Assignment) AssignmentsBuilder {
	app.list = list
	return app
}

// Now builds a new Assignments instance
func (app *assignmentsBuilder) Now() (Assignments, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Assignment in order to build an Assignments instance")
	}

	mp := map[string]Assignment{}
	for _, oneAssignment := range app.list {
		name := oneAssignment.Name()
		mp[name] = oneAssignment
	}

	if len(mp) < len(app.list) {
		diff := len(app.list) - len(mp)
		str := fmt.Sprintf("there is %d duplicate Assignments and there the Assignments instance could not be built", diff)
		return nil, errors.New(str)
	}

	return createAssignments(app.list), nil
}
