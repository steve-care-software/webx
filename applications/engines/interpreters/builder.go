package interpreters

import (
	"errors"

	creates_command "github.com/steve-care-software/syntax/applications/engines/creates/commands"
	creates_grammar "github.com/steve-care-software/syntax/applications/engines/creates/grammars"
	creates_module "github.com/steve-care-software/syntax/applications/engines/creates/modules"
	grammar_application "github.com/steve-care-software/syntax/applications/engines/grammars"
	program_application "github.com/steve-care-software/syntax/applications/engines/programs"
)

type builder struct {
	programApp program_application.Application
	grammarApp grammar_application.Application
	grammar    creates_grammar.Application
	command    creates_command.Application
	modules    creates_module.Application
	script     []byte
}

func createBuilder(
	programApp program_application.Application,
	grammarApp grammar_application.Application,
) Builder {
	out := builder{
		programApp: programApp,
		grammarApp: grammarApp,
		grammar:    nil,
		command:    nil,
		modules:    nil,
		script:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.programApp,
		app.grammarApp,
	)
}

// WithScript adds a script to the builder
func (app *builder) WithScript(script []byte) Builder {
	app.script = script
	return app
}

// WithGrammar adds a grammar create application to the builder
func (app *builder) WithGrammar(grammar creates_grammar.Application) Builder {
	app.grammar = grammar
	return app
}

// WithCommand adds a command create application to the builder
func (app *builder) WithCommand(command creates_command.Application) Builder {
	app.command = command
	return app
}

// WithModule adds a modules create application to the builder
func (app *builder) WithModules(modules creates_module.Application) Builder {
	app.modules = modules
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, []byte, error) {
	if app.grammar == nil {
		return nil, nil, errors.New("the create grammar application is mandatory in order to build an interpreter Application instance")
	}

	if app.command == nil {
		return nil, nil, errors.New("the create command application is mandatory in order to build an interpreter Application instance")
	}

	if app.modules == nil {
		return nil, nil, errors.New("the create modules application is mandatory in order to build an interpreter Application instance")
	}

	if app.script == nil {
		return nil, nil, errors.New("the script is mandatory in order to build an interpreter Application instance")
	}

	grammar, err := app.grammar.Execute()
	if err != nil {
		return nil, nil, err
	}

	tree, remaining, err := app.grammarApp.Execute(grammar, app.script)
	if err != nil {
		return nil, nil, err
	}

	command, err := app.command.Execute()
	if err != nil {
		return nil, nil, err
	}

	progIns, err := app.programApp.Execute(tree, command)
	if err != nil {
		return nil, nil, err
	}

	modules, err := app.modules.Execute()
	if err != nil {
		return nil, nil, err
	}

	return createApplication(modules, progIns), remaining, nil
}
