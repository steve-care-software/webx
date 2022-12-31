package compilers

import (
	"errors"

	"github.com/steve-care-software/webx/grammars/domain/grammars"
	"github.com/steve-care-software/webx/programs/domain/programs"
)

type executionBuilder struct {
	grammar               grammars.Grammar
	parameters            Parameters
	program               programs.Program
	pExecuteProgramModule *uint
}

func createExecutionBuilder() ExecutionBuilder {
	out := executionBuilder{
		grammar:               nil,
		parameters:            nil,
		program:               nil,
		pExecuteProgramModule: nil,
	}

	return &out
}

// Create initializes the builder
func (app *executionBuilder) Create() ExecutionBuilder {
	return createExecutionBuilder()
}

// WithGrammar adds a grammar to the builder
func (app *executionBuilder) WithGrammar(grammar grammars.Grammar) ExecutionBuilder {
	app.grammar = grammar
	return app
}

// WithParameters add parameters to the builder
func (app *executionBuilder) WithParameters(parameters Parameters) ExecutionBuilder {
	app.parameters = parameters
	return app
}

// WithProgram add a program to the builder
func (app *executionBuilder) WithProgram(program programs.Program) ExecutionBuilder {
	app.program = program
	return app
}

// WithExecuteProgramModule adds an execution program module to the builder
func (app *executionBuilder) WithExecuteProgramModule(execProgramModule uint) ExecutionBuilder {
	app.pExecuteProgramModule = &execProgramModule
	return app
}

// Now builds a new Execution instance
func (app *executionBuilder) Now() (Execution, error) {
	if app.grammar == nil {
		return nil, errors.New("the grammar is mandatory in order to build an Execution instance")
	}

	if app.parameters == nil {
		return nil, errors.New("the parameters is mandatory in order to build an Execution instance")
	}

	if app.program == nil {
		return nil, errors.New("the program is mandatory in order to build an Execution instance")
	}

	if app.pExecuteProgramModule == nil {
		return nil, errors.New("the executeProgramModule is mandatory in order to build an Execution instance")
	}

	return createExecution(app.grammar, app.parameters, app.program, *app.pExecuteProgramModule), nil
}
