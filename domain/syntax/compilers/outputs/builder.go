package outputs

import (
	"errors"

	"github.com/steve-care-software/syntax/domain/syntax/programs"
)

type builder struct {
	program programs.Program
	script  []byte
	engine  []byte
}

func createBuilder() Builder {
	out := builder{
		program: nil,
		script:  nil,
		engine:  nil,
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

// WithScript adds a script to the builder
func (app *builder) WithScript(script []byte) Builder {
	app.script = script
	return app
}

// WithEngine adds an engine to the builder
func (app *builder) WithEngine(engine []byte) Builder {
	app.engine = engine
	return app
}

// Now builds a new Output instance
func (app *builder) Now() (Output, error) {
	if app.program == nil {
		return nil, errors.New("the program is mandatory in order to build an Output instance")
	}

	if app.script != nil && len(app.script) <= 0 {
		app.script = nil
	}

	if app.engine != nil && len(app.engine) <= 0 {
		app.engine = nil
	}

	if app.script != nil && app.engine != nil {
		remaining := createRemainingWithScriptAndEngine(app.script, app.engine)
		return createOutputWithRemaining(app.program, remaining), nil
	}

	if app.script == nil {
		remaining := createRemainingWithScript(app.script)
		return createOutputWithRemaining(app.program, remaining), nil
	}

	if app.engine == nil {
		remaining := createRemainingWithEngine(app.engine)
		return createOutputWithRemaining(app.program, remaining), nil
	}

	return createOutput(app.program), nil
}
