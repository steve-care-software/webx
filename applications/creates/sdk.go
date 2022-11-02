package creates

import (
	"github.com/steve-care-software/webx/applications/creates/grammars"
	"github.com/steve-care-software/webx/applications/creates/modules"
	"github.com/steve-care-software/webx/applications/creates/selectors"
)

// NewApplication creates a new application instance
func NewApplication(
	grammar grammars.Application,
	selector selectors.Application,
	modules modules.Application,
) Application {
	return createApplication(
		grammar,
		selector,
		modules,
	)
}

// Application represents the creates application
type Application interface {
	Grammar() grammars.Application
	Selector() selectors.Application
	Modules() modules.Application
}
