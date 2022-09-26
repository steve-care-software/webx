package commands

import "errors"

type builder struct {
	commands []Command
}

func createBuilder() Builder {
	out := builder{
		commands: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList add commands to the builder
func (app *builder) WithList(commands []Command) Builder {
	app.commands = commands
	return app
}

// Now builds a new Commands instance
func (app *builder) Now() (Commands, error) {
	if app.commands != nil && len(app.commands) <= 0 {
		app.commands = nil
	}

	if app.commands == nil {
		return nil, errors.New("there must be at least 1 Command in order to build an Commands instance")
	}

	return createCommands(app.commands), nil
}
