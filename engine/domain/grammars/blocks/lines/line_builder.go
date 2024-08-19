package lines

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/webx/engine/domain/grammars/blocks/lines/tokens/elements"
)

type lineBuilder struct {
	tokens      tokens.Tokens
	execution   executions.Execution
	replacement elements.Element
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
func (app *lineBuilder) WithTokens(tokens tokens.Tokens) LineBuilder {
	app.tokens = tokens
	return app
}

// WithExecution adds execution to the builder
func (app *lineBuilder) WithExecution(execution executions.Execution) LineBuilder {
	app.execution = execution
	return app
}

// WithReplacement adds replacement to the builder
func (app *lineBuilder) WithReplacement(replacement elements.Element) LineBuilder {
	app.replacement = replacement
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.tokens == nil {
		return nil, errors.New("there must be at least 1 Token in order to build a Line instance")
	}

	if app.execution != nil && app.replacement != nil {
		return createLineWithExecutionAndReplacement(app.tokens, app.execution, app.replacement), nil
	}

	if app.execution != nil {
		return createLineWithExecution(app.tokens, app.execution), nil
	}

	if app.replacement != nil {
		return createLineWithReplacement(app.tokens, app.replacement), nil
	}

	return createLine(app.tokens), nil
}
