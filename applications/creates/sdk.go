package creates

import (
	"github.com/steve-care-software/webx/applications/creates/commands"
	"github.com/steve-care-software/webx/applications/creates/grammars"
	"github.com/steve-care-software/webx/applications/creates/modules"
)

// NewApplication creates a new application instance
func NewApplication(
	grammar grammars.Application,
	command commands.Application,
	modules modules.Application,
) Application {
	return createApplication(
		grammar,
		command,
		modules,
	)
}

// Application represents the creates application
type Application interface {
	Grammar() grammars.Application
	Command() commands.Application
	Modules() modules.Application
}
