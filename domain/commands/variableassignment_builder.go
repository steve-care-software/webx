package commands

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/bytes/criterias"
)

type variableAssignmentBuilder struct {
	assignee criterias.Criteria
	value    criterias.Criteria
}

func createVariableAssignmentBuilder() VariableAssignmentBuilder {
	out := variableAssignmentBuilder{
		assignee: nil,
		value:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *variableAssignmentBuilder) Create() VariableAssignmentBuilder {
	return createVariableAssignmentBuilder()
}

// WithAssignee adds an assignee to the builder
func (app *variableAssignmentBuilder) WithAssignee(assignee criterias.Criteria) VariableAssignmentBuilder {
	app.assignee = assignee
	return app
}

// WithValue adds a value to the builder
func (app *variableAssignmentBuilder) WithValue(value criterias.Criteria) VariableAssignmentBuilder {
	app.value = value
	return app
}

// Now builds a new VariableAssignment instance
func (app *variableAssignmentBuilder) Now() (VariableAssignment, error) {
	if app.assignee == nil {
		return nil, errors.New("the assignee is mandatory in order to build a VariableAssignment instance")
	}

	if app.value == nil {
		return nil, errors.New("the value is mandatory in order to build a VariableAssignment instance")
	}

	return createVariableAssignment(app.assignee, app.value), nil
}
