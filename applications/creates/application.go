package creates

import (
	"github.com/steve-care-software/webx/applications/creates/grammars"
	"github.com/steve-care-software/webx/applications/creates/modules"
	"github.com/steve-care-software/webx/applications/creates/selectors"
)

type application struct {
	grammar  grammars.Application
	selector selectors.Application
	modules  modules.Application
}

func createApplication(
	grammar grammars.Application,
	selector selectors.Application,
	modules modules.Application,
) Application {
	out := application{
		grammar:  grammar,
		selector: selector,
		modules:  modules,
	}

	return &out
}

// Grammar returns the grammar application
func (app *application) Grammar() grammars.Application {
	return app.grammar
}

// Command returns the selector application
func (app *application) Selector() selectors.Application {
	return app.selector
}

// Modules returns the modules application
func (app *application) Modules() modules.Application {
	return app.modules
}
