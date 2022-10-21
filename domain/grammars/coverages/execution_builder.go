package coverages

import (
	"errors"

	"github.com/steve-care-software/webx/domain/grammars"
)

type executionBuilder struct {
	expectation grammars.Suite
	result      Result
}

func createExecutionBuilder() ExecutionBuilder {
	out := executionBuilder{
		expectation: nil,
		result:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder()
}

// WithExpectation adds a suite expectation to the builder
func (app *executionBuilder) WithExpectation(expectation grammars.Suite) ExecutionBuilder {
	app.expectation = expectation
	return app
}

// WithResult adds a result to the builder
func (app *executionBuilder) WithResult(result Result) ExecutionBuilder {
	app.result = result
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.expectation == nil {
		return nil, errors.New("the suite's expectation is mandatory in order to build an Execution instance")
	}

	if app.result == nil {
		return nil, errors.New("the result is mandatory in order to build an Execution instance")
	}

	return createExecution(app.expectation, app.result), nil
}
