package processors

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens/elements"
)

type builder struct {
	execution   executions.Execution
	replacement elements.Element
}

func createBuilder() Builder {
	out := builder{
		execution:   nil,
		replacement: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithExecution adds an execution to the builder
func (app *builder) WithExecution(execution executions.Execution) Builder {
	app.execution = execution
	return app
}

// WithReplacement adds a replacement to the builder
func (app *builder) WithReplacement(replacement elements.Element) Builder {
	app.replacement = replacement
	return app
}

// WNow builds a new Processor instance
func (app *builder) Now() (Processor, error) {
	if app.execution != nil && app.replacement != nil {
		return nil, errors.New("the execution and replacement cannot both be defined in order to build a Processor instance")
	}

	if app.execution != nil {
		return createProcessorWithExecution(app.execution), nil
	}

	if app.replacement != nil {
		return createProcessorWithReplacement(app.replacement), nil
	}

	return nil, errors.New("the Processor is invalid")
}
