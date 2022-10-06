package creates

import (
	"github.com/steve-care-software/syntax/applications/engines/creates/commands"
	"github.com/steve-care-software/syntax/applications/engines/creates/grammars"
	"github.com/steve-care-software/syntax/applications/engines/creates/modules"
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
