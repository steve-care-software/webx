package commands

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/bytes/criterias"
)

type executionBuilder struct {
	application criterias.Criteria
	assignee    criterias.Criteria
}

func createExecutionBuilder() ExecutionBuilder {
	out := executionBuilder{
		application: nil,
		assignee:    nil,
	}

	return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder()
}

// WithApplication adds an application to the builder
func (app *executionBuilder) WithApplication(application criterias.Criteria) ExecutionBuilder {
	app.application = application
	return app
}

// WithAssignee adds an assignee to the builder
func (app *executionBuilder) WithAssignee(assignee criterias.Criteria) ExecutionBuilder {
	app.assignee = assignee
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.application == nil {
		return nil, errors.New("the application is mandatory in order to build an Execution instance")
	}

	if app.assignee != nil {
		return createExecutionWithAssignee(app.application, app.assignee), nil
	}

	return createExecution(app.application), nil
}
