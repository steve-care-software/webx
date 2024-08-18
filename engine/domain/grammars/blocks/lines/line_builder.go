package lines

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/replacements"
)

type lineBuilder struct {
	tokens      []string
	execution   executions.Execution
	replacement replacements.Replacement
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		tokens:      nil,
		execution:   nil,
		replacement: nil,
	}

	return &out
}

// Create initializes the builder
func (app *lineBuilder) Create() LineBuilder {
	return createLineBuilder()
}

// WithTokens add tokens to the builder
func (app *lineBuilder) WithTokens(tokens []string) LineBuilder {
	app.tokens = tokens
	return app
}

// WithExecution adds execution to the builder
func (app *lineBuilder) WithExecution(execution executions.Execution) LineBuilder {
	app.execution = execution
	return app
}

// WithReplacement adds replacement to the builder
func (app *lineBuilder) WithReplacement(replacement replacements.Replacement) LineBuilder {
	app.replacement = replacement
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.tokens != nil && len(app.tokens) <= 0 {
		app.tokens = nil
	}

	if app.tokens == nil {
		return nil, errors.New("there must be at least 1 Token in order to build a Line instance")
	}

	if app.execution != nil && app.replacement != nil {

	}

	if app.execution != nil {

	}

	if app.replacement != nil {

	}

	return nil, nil
}
