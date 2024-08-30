package lines

import (
	"errors"

	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/executions"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/processors"
	"github.com/steve-care-software/webx/engine/domain/programs/grammars/blocks/lines/tokens"
)

type lineBuilder struct {
	tokens    tokens.Tokens
	processor processors.Processor
	syscall   executions.Execution
}

func createLineBuilder() LineBuilder {
	out := lineBuilder{
		tokens:    nil,
		processor: nil,
		syscall:   nil,
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

// WithProcessor adds a processor to the builder
func (app *lineBuilder) WithProcessor(processor processors.Processor) LineBuilder {
	app.processor = processor
	return app
}

// WithSyscall adds a syscall to the builder
func (app *lineBuilder) WithSyscall(syscall executions.Execution) LineBuilder {
	app.syscall = syscall
	return app
}

// Now builds a new Line instance
func (app *lineBuilder) Now() (Line, error) {
	if app.tokens == nil {
		return nil, errors.New("there must be at least 1 Token in order to build a Line instance")
	}

	if app.processor != nil && app.syscall != nil {
		return createLineWithProcessorAndSyscall(app.tokens, app.processor, app.syscall), nil
	}

	if app.processor != nil {
		return createLineWithProcessor(app.tokens, app.processor), nil
	}

	if app.syscall != nil {
		return createLineWithSyscall(app.tokens, app.syscall), nil
	}

	return createLine(app.tokens), nil
}
