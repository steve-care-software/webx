package coverages

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/bytes/grammars"
)

type executionBuilder struct {
	suite grammars.Suite
	line  Line
}

func createExecutionBuilder() ExecutionBuilder {
	out := executionBuilder{
		suite: nil,
		line:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder()
}

// WithSuite adds a suite to the builder
func (app *executionBuilder) WithSuite(suite grammars.Suite) ExecutionBuilder {
	app.suite = suite
	return app
}

// WithLine adds a line to the builder
func (app *executionBuilder) WithLine(line Line) ExecutionBuilder {
	app.line = line
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.suite == nil {
		return nil, errors.New("the suite is mandatory in order to build an Execution instance")
	}

	if app.line == nil {
		return nil, errors.New("the line is mandatory in order to build an Execution instance")
	}

	return createExecution(app.suite, app.line), nil
}
