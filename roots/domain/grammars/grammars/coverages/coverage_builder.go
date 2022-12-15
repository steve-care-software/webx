package coverages

import (
	"errors"

	"github.com/steve-care-software/webx/roots/domain/grammars/grammars"
)

type coverageBuilder struct {
	token      grammars.Token
	executions Executions
}

func createCoverageBuilder() CoverageBuilder {
	out := coverageBuilder{
		token:      nil,
		executions: nil,
	}

	return &out
}

// Create initializes the builder
func (app *coverageBuilder) Create() CoverageBuilder {
	return createCoverageBuilder()
}

// WithToken adds a token to the builder
func (app *coverageBuilder) WithToken(token grammars.Token) CoverageBuilder {
	app.token = token
	return app
}

// WithExecutions add executions to the builder
func (app *coverageBuilder) WithExecutions(executions Executions) CoverageBuilder {
	app.executions = executions
	return app
}

// Now builds a new Coverage instance
func (app *coverageBuilder) Now() (Coverage, error) {
	if app.token == nil {
		return nil, errors.New("the token is mandatory in order to build a Coverage instance")
	}

	if app.executions == nil {
		return nil, errors.New("the executions is mandatory in order to build a Coverage instance")
	}

	return createCoverage(app.token, app.executions), nil
}
