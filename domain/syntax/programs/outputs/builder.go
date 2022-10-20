package outputs

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/programs"
)

type builder struct {
	program   programs.Program
	remaining []byte
}

func createBuilder() Builder {
	out := builder{
		program:   nil,
		remaining: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithProgram adds a program to the builder
func (app *builder) WithProgram(program programs.Program) Builder {
	app.program = program
	return app
}

// WithRemaining adds a remaining to the builder
func (app *builder) WithRemaining(remaining []byte) Builder {
	app.remaining = remaining
	return app
}

// Now builds a new Output instance
func (app *builder) Now() (Output, error) {
	if app.program == nil {
		return nil, errors.New("the program is mandatory in order to build an Output instance")
	}

	if app.remaining != nil && len(app.remaining) <= 0 {
		app.remaining = nil
	}

	if app.remaining != nil {
		return createOutputWithRemaining(app.program, app.remaining), nil
	}

	return createOutput(app.program), nil
}
