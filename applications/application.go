package applications

import (
	"github.com/steve-care-software/syntax/applications/actions"
	"github.com/steve-care-software/syntax/applications/languages"
	"github.com/steve-care-software/syntax/applications/modules"
)

type application struct {
	action   actions.Application
	language languages.Application
	module   modules.Application
}

func createApplication(
	action actions.Application,
	language languages.Application,
	module modules.Application,
) Application {
	out := application{
		action:   action,
		language: language,
		module:   module,
	}

	return &out
}

// Action returns the action application
func (app *application) Action() actions.Application {
	return app.action
}

// Language returns the language application
func (app *application) Language() languages.Application {
	return app.language
}

// Module returns the module application
func (app *application) Module() modules.Application {
	return app.module
}
