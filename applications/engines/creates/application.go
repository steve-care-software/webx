package creates

import (
	"github.com/steve-care-software/syntax/applications/engines/creates/commands"
	"github.com/steve-care-software/syntax/applications/engines/creates/grammars"
	"github.com/steve-care-software/syntax/applications/engines/creates/modules"
)

type application struct {
	grammar grammars.Application
	command commands.Application
	modules modules.Application
}

func createApplication(
	grammar grammars.Application,
	command commands.Application,
	modules modules.Application,
) Application {
	out := application{
		grammar: grammar,
		command: command,
		modules: modules,
	}

	return &out
}

// Grammar returns the grammar application
func (app *application) Grammar() grammars.Application {
	return app.grammar
}

// Command returns the command application
func (app *application) Command() commands.Application {
	return app.command
}

// Modules returns the modules application
func (app *application) Modules() modules.Application {
	return app.modules
}
